# FTL examples

This repository contains an [FTL](https://github.com/block/ftl) implementation of
GCP's [microservices demo](https://github.com/GoogleCloudPlatform/microservices-demo), Online Boutique.

The example highlights how to use FTL to build an online store backend, and use its schema to automatically generate
clients for a mobile (Flutter) app and React frontend.

> [!NOTE]
> If using Hermit, be sure to open your editor from the activated environment in
> order to get the correct environment variables.

## Running

In order to have all the tools you need to build and run the examples, either:

1. Use [Hermit](https://github.com/cashapp/hermit): `. ./bin/activate-hermit`
2. Use Homebrew: `brew tap block/ftl && brew install flutter ftl go just node`

> [!WARNING]
> Be aware that Homebrew will install the most recent versions of these dependencies, which may not be compatible with
> the examples. While we do try to keep this repo up to date, if you run into issues, please use Hermit.

### Preparation

To see a list of commands that can be run, run `just --list` in the root of the
project.

### Start FTL, code gen, and backend services with hot-reloading

```bash
just dev
```

### Start the React web app

```bash
just web
```

### Flutter app

Flutter is better run via an [IDE](https://docs.flutter.dev/get-started/editor?tab=vscode) like VSCode or Android Studio.

### Interact with the boutique

Add a new item to your cart using the command-line:

```bash
curl -i --json '{"userId": "Larry", "item": {"productId": "OLJCESPC7Z", "quantity": 1}}' localhost:8891/cart/add
```

Response

```
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Vary: Origin
Date: Wed, 14 Feb 2024 03:00:03 GMT
Content-Length: 2

{}
```

Checkout the cart:

```bash
curl --json '{
  "userCurrency": "AUD",
  "address": {
    "streetAddress": "1600 Amphitheatre Parkway",
    "city": "Mountain View",
    "state": "CA",
    "country": "USA",
    "zipCode": 94043
  },
  "email": "alec@tbd.email",
  "creditCard": {
    "number": "423412341234",
    "cvv": 123,
    "expirationYear": 2029,
    "expirationMonth": 6
  }
}
' localhost:8891/checkout/Larry | jq .
```

Response

```
{
  "id": "19031e2c-a4b3-49fe-bbb8-464bc8707559",
  "shippingAddress": {
    "city": "Mountain View",
    "country": "USA",
    "state": "CA",
    "streetAddress": "1600 Amphitheatre Parkway",
    "zipCode": 94043
  },
  "shippingCost": {
    "currencyCode": "AUD",
    "nanos": 387449163,
    "units": 11
  },
  "shippingTrackingId": "LS-44971-224078566"
}
```

Open up the [FTL console](http://localhost:8892). This will let you navigate the
logs, traces and structure of the FTL modules running locally. In the timeline,
if you click on one of the calls that we just issued you should see a call trace.

Open up your editor, alter the data structures, alter the logic, and see how it
is all reflected in FTL automatically.
