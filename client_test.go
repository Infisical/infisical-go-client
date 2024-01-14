package infisicalclient

//func TestNewClient(t *testing.T) {
//	serviceToken := "xxxx"
//	client, err := NewClient(Config{HostURL: "https://app.infisical.com", ServiceToken: serviceToken})
//	if err != nil {
//		t.Error(err)
//	}
//	response, err := client.CallGetServiceTokenDetailsV2()
//
//	if err != nil {
//		t.Error(err)
//	}
//
//	println(response.Workspace)
//
//	secrets, err := client.FetchAndDecodeSecrets("6500618f61d12cd3d808036a", "dev")
//
//	for k, v := range secrets {
//		println(k + ":" + v)
//	}
//}
