# passkey-impl
This repository is a sample implementation of passkey using React and Golang. A server side uses [go-webauthn/webauthn](https://github.com/go-webauthn/webauthn).

<video src="https://github.com/ksrnnb/passkey-impl/assets/48155865/76758daa-755a-4db9-b3c3-9f77cd4789c2"></video>

# How to set up
## set up server side
```bash
cd server
go run main.go
```

## set up client side
```bash
cd client
npm run start
```

# How to use
## login with user id
you can login with user id `sample`. password is unnecessary in this application.

## register your passkey
After login, you can register passkey.

## login with passkey
After register passkey, you can login with passkey using auto fill. Auto fill screen will appear when you click user id form.
