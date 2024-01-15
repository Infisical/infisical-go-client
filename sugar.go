package infisicalclient

/*
*
Get all secrets for a given environment and workspace, with E2EE supported
*/
func (client Client) FetchAndDecodeSecrets(workspaceId string, environment string) (map[string]string, error) {
	var secrets map[string]string
	secrets = make(map[string]string)

	response, err := client.CallGetServiceTokenDetailsV2()

	if err != nil {
		return nil, err
	}

	projectKey, err := DecryptProjectKey(client.cnf.ServiceToken, response)

	res, err := client.CallGetSecretsV3(GetEncryptedSecretsV3Request{Environment: environment, WorkspaceId: workspaceId})

	if err != nil {
		return nil, err
	}

	for _, secret := range res.Secrets {
		key, value, err := DecryptSecret(secret, projectKey)

		if err != nil {
			return nil, err
		}

		secrets[key] = value
	}

	return secrets, nil
}
