package infisicalclient

import "fmt"

func (client Client) CallGetServiceTokenDetailsV2() (GetServiceTokenDetailsResponse, error) {
	var tokenDetailsResponse GetServiceTokenDetailsResponse
	response, err := client.cnf.HttpClient.
		R().
		SetResult(&tokenDetailsResponse).
		Get("api/v2/service-token")

	if err != nil {
		return GetServiceTokenDetailsResponse{}, fmt.Errorf("CallGetServiceTokenDetails: Unable to complete api request [err=%s]", err)
	}

	if response.IsError() {
		return GetServiceTokenDetailsResponse{}, fmt.Errorf("CallGetServiceTokenDetails: Unsuccessful response: [response=%s]", response)
	}

	return tokenDetailsResponse, nil
}

func (client Client) CallGetSecretsV3(request GetEncryptedSecretsV3Request) (GetEncryptedSecretsV3Response, error) {
	var secretsResponse GetEncryptedSecretsV3Response
	requestToBeMade := client.cnf.HttpClient.
		R().
		SetResult(&secretsResponse).
		SetHeader("Authorization", "Bearer "+client.cnf.ServiceToken).
		SetQueryParam("environment", request.Environment).
		SetQueryParam("workspaceId", request.WorkspaceId)

	if request.SecretPath != "" {
		requestToBeMade.SetQueryParam("secretsPath", request.SecretPath)
	}
	if !request.IncludeImports {
		requestToBeMade.SetQueryParam("include_imports", "true")
	}

	response, err := requestToBeMade.
		Get("api/v3/secrets")

	if err != nil {
		return GetEncryptedSecretsV3Response{}, fmt.Errorf("CallGetSecretsV2: Unable to complete api request [err=%s]", err)
	}

	if response.IsError() {
		return GetEncryptedSecretsV3Response{}, fmt.Errorf("CallGetSecretsV2: Unsuccessful response: [response=%v]", response.RawResponse)
	}

	return secretsResponse, nil
}
