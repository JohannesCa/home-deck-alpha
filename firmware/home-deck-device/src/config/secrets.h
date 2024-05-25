#ifndef SECRETS_H
#define SECRETS_H

const char WIFI_SSID[] = "Starlink";
const char WIFI_PASSWORD[] = "@sk7zulu";

#define THINGNAME "esp8266-opener-alpha"

int8_t TIME_ZONE = -3;

const char MQTT_HOST[] = "a3qkll921pt2mi-ats.iot.us-east-1.amazonaws.com";


static const char cacert[] PROGMEM = R"EOF(
-----BEGIN CERTIFICATE-----
MIIGAjCCBOqgAwIBAgIQBX2O4/fGrI8FrUCuUvBtVjANBgkqhkiG9w0BAQsFADA8
MQswCQYDVQQGEwJVUzEPMA0GA1UEChMGQW1hem9uMRwwGgYDVQQDExNBbWF6b24g
UlNBIDIwNDggTTAxMB4XDTIzMTIxNDAwMDAwMFoXDTI0MTIwNTIzNTk1OVowKDEm
MCQGA1UEAwwdKi5pb3QudXMtZWFzdC0xLmFtYXpvbmF3cy5jb20wggEiMA0GCSqG
SIb3DQEBAQUAA4IBDwAwggEKAoIBAQC0k2ft+pRoOgh6BZnZAK+1yQAiq5Ge9Euz
t1YtUYHZYx7ucGoK15XA9F5xsQHH5pqa70eECE6XhmnQRfVlPLw5Hb41W2bFeJj3
seHx3bi7TqRNhx9gF6KKrLGo4d+eJzIAdTrmXTjCr+mMLlE4Ep01eZRWbZDcSx1U
lRZIrWU/10xQyUl8evOByz4k0GYATVFKP0BfVxuL1DNTsRlqzjHC8WPAGq7z/Wbi
KZ+DXqGmeUGmvdQx5RvFqC3OCjRG3SSzC0vcpY3dEtNcdbTSceMdU6FVJjQPSgwx
gB1j1y6uhCWO2LaBK2xBtvKuBxFaL8BbaHRg98HRd0vXRcUW4rkzAgMBAAGjggMS
MIIDDjAfBgNVHSMEGDAWgBSBuA5jiokSGOX6OztQlZ/m5ZAThTAdBgNVHQ4EFgQU
caZbCxMNhVZ7uob+cE5DaiWBIE8wRQYDVR0RBD4wPIIbaW90LnVzLWVhc3QtMS5h
bWF6b25hd3MuY29tgh0qLmlvdC51cy1lYXN0LTEuYW1hem9uYXdzLmNvbTATBgNV
HSAEDDAKMAgGBmeBDAECATAOBgNVHQ8BAf8EBAMCBaAwHQYDVR0lBBYwFAYIKwYB
BQUHAwEGCCsGAQUFBwMCMDsGA1UdHwQ0MDIwMKAuoCyGKmh0dHA6Ly9jcmwucjJt
MDEuYW1hem9udHJ1c3QuY29tL3IybTAxLmNybDB1BggrBgEFBQcBAQRpMGcwLQYI
KwYBBQUHMAGGIWh0dHA6Ly9vY3NwLnIybTAxLmFtYXpvbnRydXN0LmNvbTA2Bggr
BgEFBQcwAoYqaHR0cDovL2NydC5yMm0wMS5hbWF6b250cnVzdC5jb20vcjJtMDEu
Y2VyMAwGA1UdEwEB/wQCMAAwggF9BgorBgEEAdZ5AgQCBIIBbQSCAWkBZwB1AO7N
0GTV2xrOxVy3nbTNE6Iyh0Z8vOzew1FIWUZxH7WbAAABjGnJk6gAAAQDAEYwRAIg
AqYOtxbFag+jsRSnX/dVyPYh3qteOpCn4eUXkTUX7wICIHGOfZlQUBNuNCcjsT0k
13AL6OjiPe3j7WariS684kJJAHUASLDja9qmRzQP5WoC+p0w6xxSActW3SyB2bu/
qznYhHMAAAGMacmTzgAABAMARjBEAiAoUVfDJoMe8gK/rYTvjj4fHNPrrAnO8Mhy
1N6iGK6LngIgWDj6X03ESCYWTyyLhOXYgpLwUpcl38WkOi/uzE6nxyIAdwDatr9r
P7W2Ip+bwrtca+hwkXFsu1GEhTS9pD0wSNf7qwAAAYxpyZO+AAAEAwBIMEYCIQC7
2UGBh6xrR1DM+MwbNwUmj+qaRwoZI6lVKGDz/azxoQIhAPUuVuuVTryPqzNoyUsw
BVcP7GchjnjHo94vSzSkrOLHMA0GCSqGSIb3DQEBCwUAA4IBAQBV3uYeBy/SFjIo
jSSno9bXqwEgbMI1drei/Z+3rC9gLxL3Vm/y38hDwfi6e/LETClQh0nmZJsfo4pR
VND6cxwRTHgZhwA52xX0Filln2UrdaXvy+OfYeszTBsPsVO2GLqO83D28M+Tz8B4
yAILzjm/h7P3Y8abY8n+tsZgNVkcuSuOzBvGKUg2eJj8DAcPys6nsUcc3K4uJikO
7Cp6e7y/qs7Ok4lcuPhF+lxphmkSi7fxCke7/UhuMyT18IwFFadxyOq66r7FC3OI
XNyLhcQ83lDeSX9//lxpLpbXsynPPDEH0ne3koVKJUaxq8GlUlmmdDv6SK72cnj7
zH1a4EjV
-----END CERTIFICATE-----
)EOF";


// Copy contents from XXXXXXXX-certificate.pem.crt here ▼
static const char client_cert[] PROGMEM = R"EOF(
-----BEGIN CERTIFICATE-----
MIIDWTCCAkGgAwIBAgIUDR4jaZ8kHyU2yUsHAtw1T7xy8b4wDQYJKoZIhvcNAQEL
BQAwTTFLMEkGA1UECwxCQW1hem9uIFdlYiBTZXJ2aWNlcyBPPUFtYXpvbi5jb20g
SW5jLiBMPVNlYXR0bGUgU1Q9V2FzaGluZ3RvbiBDPVVTMB4XDTI0MDUwOTAyMDkw
MloXDTQ5MTIzMTIzNTk1OVowHjEcMBoGA1UEAwwTQVdTIElvVCBDZXJ0aWZpY2F0
ZTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKeeZ3HDN17HwV8+ztNp
FkccHXBijxZ8IwHaqTZqvfSsAn9utPrHCaLitoZ9qaHnDyCKgSBdsJ3kCf6k5Uhr
uaPjFVPy2DADpSahOIaIhYV2fPnds+ZgW4t3FMOttS65r45YhPEXFAIRD4m1U5CV
A6iCBIONIprpImzsKYjNVpm6v69Q8AGlzFVbBU0UcmCdyVDsWLF6H8wTBkjmDWiQ
5FtQWixc4VZIjQhz+aKDPMrcqarucgPBnyWI7dIgB+iFB8KVaBy0aBsKcuqPoLNB
hoKAJMdCyJjs3O6luN884Ykt/2RIa3XqC4bgF4zbCTAWcEgeotDL1Pf8UK2ZHHR5
bv8CAwEAAaNgMF4wHwYDVR0jBBgwFoAUBEmHKvSGsFBdxNHhihTXZd+l0okwHQYD
VR0OBBYEFDnCcUNOg0iIirgyrEw0hBsEo0I4MAwGA1UdEwEB/wQCMAAwDgYDVR0P
AQH/BAQDAgeAMA0GCSqGSIb3DQEBCwUAA4IBAQAf4YvyTmatXLoVeKNM9m9S9Ku1
KhNoqAqAv5o21NefhFmJIs6lpztoLyQ9a3kHKpOgFCSCWYhoI4RLw/UytsBAJ/sD
j/piX3hFxfgK9wxeQPGKr8BwVELxrYUgFcmd09MWmuZAVBif5NGMwjZW5Pjj1XfG
//28cCDVrCsmoyAWrodlD7n01r5Xil1hE64SEZOgkCSel/W0f+PPivrQfFSNa0X2
x7xQy1VaLXVstMuFNSjyiwaIB3oH8TPK6TGzLIdugJt9E7GFDQ4hmu6DK+STAJLK
h76x1chEAm/aAM3fusMDVo4viEIQ8nhxJNbwCoJjArr8lngc7Nv5Xne/plB0
-----END CERTIFICATE-----
)EOF";


// Copy contents from  XXXXXXXX-private.pem.key here ▼
static const char privkey[] PROGMEM = R"EOF(
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAp55nccM3XsfBXz7O02kWRxwdcGKPFnwjAdqpNmq99KwCf260
+scJouK2hn2poecPIIqBIF2wneQJ/qTlSGu5o+MVU/LYMAOlJqE4hoiFhXZ8+d2z
5mBbi3cUw621LrmvjliE8RcUAhEPibVTkJUDqIIEg40imukibOwpiM1Wmbq/r1Dw
AaXMVVsFTRRyYJ3JUOxYsXofzBMGSOYNaJDkW1BaLFzhVkiNCHP5ooM8ytypqu5y
A8GfJYjt0iAH6IUHwpVoHLRoGwpy6o+gs0GGgoAkx0LImOzc7qW43zzhiS3/ZEhr
deoLhuAXjNsJMBZwSB6i0MvU9/xQrZkcdHlu/wIDAQABAoIBAAstQtSRIdYOF9TU
CVKLhVLEiyn9SHFFYcgz3mikgXoZJil9JlA6iWf1Ghqpsm/7SkTvN2EDHk3I8Omr
8SX9GOnWsItifAvnklGbGnvOCqk1l2vJ3r8T2cbY7+eEY7pz40LooCS7Vzg0Dwct
sh7H8PCBiu+Cr7YIqaT9sl9O4eFqOMJEZ50227cypLECPsdOC+Bk9EwPvO2KVkKi
3RZD3VDLJbwiQXCdFCVDtR+I2jX7c3c7MOa62w3Id9Zd0r8uUaQItbFV1UigS7O8
ySud2HCYuYpw0uxuuBt4rUrrEqlB9qAq7FtZBi23RHcP2wOI99pxOK7RLEz3Li/L
kSUw1dECgYEA3U4w/EWD53BwhzIj2WBhjsSkTnTvDnB1skKEn3vue/ffzy0j7+i/
8w3XwT4RBeLwsu23an+2inT0qfpKMuqjVJ2Qq7BX19k7EyY14+zDV4Hok1Miiwb3
hHYAdudCW5hzrc9xvaT1AI7g8ieNvJPPl3gqu4sD/tign2RjjIpfs4MCgYEAweWR
QVGaole21JdtvVj/ACzcy7q5o/hrx7l0xwTUz+voIoe3nuUCSe/XA1y3APZaPv0H
sB4CTmnXTd7EizzaRTm8WHsm3eisjr9IshcKinr1aELldlICM6vMslnnr0eCNGZ3
+yCp+C3sgbNQOVB5Qzr57hDU5KR8MjmGHWb/MdUCgYB12bHUZeB8WzBRjq3wkJbj
RGyLK1IyykijCTE/DOxPMA0Lcup0U+oEMRb4XWkLTXMIlfbay/AVg3rXGeZxfupD
oTTteRcOZyxCYlLIcckEzLVh3x/CEY+2BnDw4O5BgMn5gUL8zt950JyR1+Zuyff+
nYuEG9+g0xOWTguFSBr2qQKBgAmAAPTtVaE8oPWbkcHkOWI90Hee53eUNVusx370
yqSBT3dZpPKBD/KlcKh2g3TlnRhW92h2n/BNmUH0dzbHPX6EzKidKBNZUhMQZjRr
kpWkXp/EH1ADxp3XfgkDBKn+2+9LUfeY6xsAzQSxz21vdvGHUzWJ15R4RqmDTgHk
yP4dAoGBALmmpZ0O2xg/67kl8EZouoQxGTQ+3lC39qZkQZ3q36vInkpDwqvA0afp
UbKtLp32CuEkARXPN+qxiEMsXSwOLpaqGOtG5oEAVxGIgqyykVV0gWug2IPD/6rB
EjBOf+H4tw0FIqmjKGW6S+zDksXPDSsVCZigEGOObh2gDWJqj0nI
-----END RSA PRIVATE KEY-----
)EOF";


#endif // SECRETS_H