module github.com/rajatjindal/gitops-tailscale-authkey-action

go 1.24.2

require (
	github.com/spinframework/spin-go-sdk/v2 v2.0.0-00010101000000-000000000000
	github.com/ydnar/wasi-http-go v0.0.0-20250324053847-ca78b3198aeb
	golang.org/x/oauth2 v0.30.0
	tailscale.com/client/tailscale/v2 v2.0.0-20250509161557-5fad10cf3a33
)

require (
	github.com/tailscale/hujson v0.0.0-20220506213045-af5ed07155e5 // indirect
	go.bytecodealliance.org/cm v0.2.2 // indirect
)

replace github.com/spinframework/spin-go-sdk/v2 => github.com/spinframework/spin-go-sdk/v2 v2.0.0-20250422162322-8ffe6d3efa29
