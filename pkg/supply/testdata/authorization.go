package testdata

const EnvWithAuthorization = `{
"authorization": [
{
  "binding_guid": "cf5fbc53-f243-4411-a1d3-f6e254ec7bef",
  "binding_name": null,
  "credentials": {
  "object_store": {
  "access_key_id": "myawstestaccesskeyid",
  "bucket": "my-bucket",
  "host": "s3-eu-central-1.amazonaws.com",
  "region": "eu-central-1",
  "secret_access_key": "mysecretaccesskey",
  "uri": "s3://myawstestaccesskeyid:TTe84CgHEQ%mysecretaccesskey@s3-eu-central-1.amazonaws.com/my-bucket",
  "username": "my-username"
  },
  "ui_url": "https://4b0c2b7a-1279-4352-aasdfsadf--asdfasdf-ams-ui.authorization-dev.cfapps.ls.domain",
  "url": "https://ams.cert.cfapps.ls.domain/",
  "value_help_certificate_issuer": "{\"Country\":[\"DE\"],\"Organization\":[\"SAP SE\"],\"Locality\":[\"Walldorf\"],\"CommonName\":\"SAP Cloud Root CA\"}",
  "value_help_certificate_subject": "{\"Country\":[\"DE\"],\"Organization\":[\"SAP SE\"],\"OrganizationalUnit\":[\"SAP Cloud Platform Clients\",\"Canary\"],\"Locality\":[\"AMS\"],\"CommonName\":\"ValueHelpmTLSCert\"}"
  },
  "instance_guid": "my-instance-guid",
  "instance_name": "my-ams-instance",
  "label": "authorization",
  "name": "my-ams-instance",
  "plan": "application",
  "provider": null,
  "syslog_drain_url": null,
  "tags": [],
  "volume_mounts": []
}
]
}`

const EnvWithAuthorizationDev = `{
"authorization-dev": [
{
  "binding_guid": "cf5fbc53-f243-4411-a1d3-f6e254ec7bef",
  "binding_name": null,
  "credentials": {
  "object_store": {
  "access_key_id": "myawstestaccesskeyid",
  "bucket": "my-bucket",
  "host": "s3-eu-central-1.amazonaws.com",
  "region": "eu-central-1",
  "secret_access_key": "mysecretaccesskey",
  "uri": "s3://myawstestaccesskeyid:TTe84CgHEQ%mysecretaccesskey@s3-eu-central-1.amazonaws.com/my-bucket",
  "username": "my-username"
  },
  "ui_url": "https://4b0c2b7a-1279-4352-aasdfsadf--asdfasdf-ams-ui.authorization-dev.cfapps.ls.domain",
  "url": "https://ams.cert.cfapps.ls.domain/",
  "value_help_certificate_issuer": "{\"Country\":[\"DE\"],\"Organization\":[\"SAP SE\"],\"Locality\":[\"Walldorf\"],\"CommonName\":\"SAP Cloud Root CA\"}",
  "value_help_certificate_subject": "{\"Country\":[\"DE\"],\"Organization\":[\"SAP SE\"],\"OrganizationalUnit\":[\"SAP Cloud Platform Clients\",\"Canary\"],\"Locality\":[\"AMS\"],\"CommonName\":\"ValueHelpmTLSCert\"}"
  },
  "instance_guid": "my-instance-guid",
  "instance_name": "my-ams-instance",
  "label": "authorization-dev",
  "name": "my-ams-instance",
  "plan": "application",
  "provider": null,
  "syslog_drain_url": null,
  "tags": [],
  "volume_mounts": []
}
]
}`

const EnvWithUserProvidedAuthorization = `{
"user-provided": [
{
  "binding_guid": "cf5fbc53-f243-4411-a1d3-f6e254ec7bef",
  "binding_name": null,
  "credentials": {
  "object_store": {
  "access_key_id": "myawstestaccesskeyid",
  "bucket": "my-bucket",
  "host": "s3-eu-central-1.amazonaws.com",
  "region": "eu-central-1",
  "secret_access_key": "mysecretaccesskey",
  "uri": "s3://myawstestaccesskeyid:TTe84CgHEQ%mysecretaccesskey@s3-eu-central-1.amazonaws.com/my-bucket",
  "username": "my-username"
  },
  "ui_url": "https://4b0c2b7a-1279-4352-aasdfsadf--asdfasdf-ams-ui.authorization-dev.cfapps.ls.domain",
  "url": "https://ams.cert.cfapps.ls.domain/",
  "value_help_certificate_issuer": "{\"Country\":[\"DE\"],\"Organization\":[\"SAP SE\"],\"Locality\":[\"Walldorf\"],\"CommonName\":\"SAP Cloud Root CA\"}",
  "value_help_certificate_subject": "{\"Country\":[\"DE\"],\"Organization\":[\"SAP SE\"],\"OrganizationalUnit\":[\"SAP Cloud Platform Clients\",\"Canary\"],\"Locality\":[\"AMS\"],\"CommonName\":\"ValueHelpmTLSCert\"}"
  },
  "instance_guid": "my-instance-guid",
  "instance_name": "my-ams-instance",
  "label": "authorization-dev",
  "name": "my-ams-instance",
  "plan": "application",
  "provider": null,
  "syslog_drain_url": null,
  "tags": ["authorization"],
  "volume_mounts": []
}
]
}`

const EnvWithIASAuthWithClientSecret = `{
"identity": [
	{
	  "binding_guid": "cf5fbc53-f243-4411-a1d3-f6e254ec7bef",
	  "binding_name": null,
		"credentials": {
			"authorization_object_store": {
				"access_key_id": "myawstestaccesskeyid",
				"bucket": "my-bucket",
				"host": "s3-eu-central-1.amazonaws.com",
				"region": "eu-central-1",
				"secret_access_key": "mysecretaccesskey",
				"uri": "s3://myawstestaccesskeyid:mysecretaccesskey@my-bucket-from-ias-svc",
				"username": "my-username"
			},
			"authorization_url": "https://ams.url.from.identity/svc",
			"authorization_ui_url": "https://4b0c2b7a-1279-4352-aasdfsadf--asdfasdf-ams-ui.authorization-dev.cfapps.ls.domain",
			"authorization_value_help_certificate_issuer": "{\"Country\":[\"DE\"],\"Organization\":[\"SAP SE\"],\"Locality\":[\"Walldorf\"],\"CommonName\":\"SAP Cloud Root CA\"}",
			"authorization_value_help_certificate_subject": "{\"Country\":[\"DE\"],\"Organization\":[\"SAP SE\"],\"OrganizationalUnit\":[\"SAP Cloud Platform Clients\",\"Canary\",\"someuid\"],\"Locality\":[\"AMS\"],\"CommonName\":\"ValueHelpmTLSCert\"}",
			"clientid": "601AED1D-0063-4B5B-8FCE-06FC75AF957C",
			"domain": "accounts400.ondemand.com",
			"domains": [
				"accounts400.ondemand.com"
			],
			"clientSecret": "mysecret",
			"url": "https://mytenant.accounts400.ondemand.com",
			"zone_uuid": "3276660F-1C15-4311-9A29-CE8704006361"
		},
		"instance_guid": "623D51BC-F961-4D6E-B737-0FC08B495755",
		"instance_name": "ias-ams-node",
		"label": "identity",
		"name": "ias-ams-node",
		"plan": "application",
		"provider": null,
		"syslog_drain_url": null,
		"tags": [],
		"volume_mounts": []
	}
]
}`

const EnvWithIASAuthX509 = `{
"identity": [
	{
	  "binding_guid": "cf5fbc53-f243-4411-a1d3-f6e254ec7bef",
	  "binding_name": null,
		"credentials": {
			"authorization_object_store": {
				"access_key_id": "myawstestaccesskeyid",
				"bucket": "my-bucket",
				"host": "s3-eu-central-1.amazonaws.com",
				"region": "eu-central-1",
				"secret_access_key": "mysecretaccesskey",
				"uri": "s3://myawstestaccesskeyid:mysecretaccesskey@my-bucket-from-ias-svc",
				"username": "my-username"
			},
			"authorization_url": "https://ams.url.from.identity/svc",
			"authorization_ui_url": "https://4b0c2b7a-1279-4352-aasdfsadf--asdfasdf-ams-ui.authorization-dev.cfapps.ls.domain",
			"authorization_value_help_certificate_issuer": "{\"Country\":[\"DE\"],\"Organization\":[\"SAP SE\"],\"Locality\":[\"Walldorf\"],\"CommonName\":\"SAP Cloud Root CA\"}",
			"authorization_value_help_certificate_subject": "{\"Country\":[\"DE\"],\"Organization\":[\"SAP SE\"],\"OrganizationalUnit\":[\"SAP Cloud Platform Clients\",\"Canary\",\"someuid\"],\"Locality\":[\"AMS\"],\"CommonName\":\"ValueHelpmTLSCert\"}",
			"certificate": "-----BEGIN CERTIFICATE-----\nMIICrDCCAZSgAwIBAgIBfTANBgkqhkiG9w0BAQsFADASMRAwDgYDVQQKEwd0ZXN0\nT3JnMCIYDzAwMDEwMTAxMDAwMDAwWhgPMDAwMTAxMDEwMDAwMDBaMBkxFzAVBgNV\nBAoTDnRlc3RTdWJqZWN0T3JnMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKC\nAQEAz7YopBPH6RbY0taFoFSBVmMgLqWIG0aICgY6fd9ruA+PKD2Dl4IyLzudrVw5\naZipGzBlStLQWwOjPDtU7HDBV8mNLv4IgaXWutSEKNWkrVhekaoYqJsKqSt9yiBO\nVavd1Mz2SbbwQ3lajMz2L0TKTW5jrpx6RYQ6ueVgdPQsNPWyE458pQJfrIiWO9iZ\nqt+oCCTQxkJ4E776zKCEI+qt4L78bYUwYGpIMQR4To6GHvqCTwYuwUDmw/IYaunf\nzqXMMhLG4bmw50o0qrE+erwdi9RANbsyI5+k8eHgAAFhCNTDjkI4zicR7fw3C+Yy\n11o8pMv26laGVMhmwMyEd97sjwIDAQABowIwADANBgkqhkiG9w0BAQsFAAOCAQEA\nhCrgBSWkmvxKVfgCUbi1QoziJIayvvMh1OYkU7hA6PdR9FYwKYOnHqaeRFbTR9GX\nQ9Cj9vAOPjbvjxkWne8gJpNdV4faTVd68X2BeSXr6XuH7IyQYfUhNDdZOiIOxl4v\nq8Y4gGyBvX4rnWckJ1770A0FdqylQiI4SaLeA4aQ7rPu57fwtXCRqKMkzj+jpCMW\nt0qiHJt5M1cW5YnB2QO8VMBkXLMNePRvstWQweSF1GKOFVzzJL+b4kAYBIRgJ20A\nfZwfbtPhJ1wn92V3UBMf3ScXextNZ/2m6iGyOvNHDc6On6sNQuoiZCgmvwDv6PGL\ncJeKuDIDTYO78B49UvLI+A==\n-----END CERTIFICATE-----",
			"certificate_expires_at": "2021-11-28T09:57:08Z",
			"clientid": "601AED1D-0063-4B5B-8FCE-06FC75AF957C",
			"domain": "accounts400.ondemand.com",
			"domains": [
				"accounts400.ondemand.com"
			],
			"key": "-----BEGIN RSA PRIVATE KEY-----\nY/ph+LwaZeAj29WJxhWE7yV4vAw+LtH7CJjPzLKAKDdawHYKUvPMrKs44LxISjFL\nK+Ai/DDFJvxyZ/Jy3vthFOhJ3LgvLKHF23x0SEMdV51v/0lI49dn7KYflyiRcQ4L\nrH0wi2LeA7sgSyH6V7pVGOSvx2STUMMnUCDYlm/5CJEhpg7DIiNXc4kr97pizgve\nUe8/KoMtoEp3bwINQDzGV+uxgFBcOMJlc6fdbAiFlpJT3obyG1Cw+uxSLCiJuJao\nRkt08lMfsWMpSAX6pHp0kzoKpbJy68NCgBYrv2oji8eIkRrEbcSC2OSh/UfqecKU\n4EQ6j6qGtwnifRm9P4KFgQ==\n-----END RSA PRIVATE KEY-----",
			"url": "https://mytenant.accounts400.ondemand.com",
			"zone_uuid": "3276660F-1C15-4311-9A29-CE8704006361"
		},
		"instance_guid": "623D51BC-F961-4D6E-B737-0FC08B495755",
		"instance_name": "ias-ams-node",
		"label": "identity",
		"name": "ias-ams-node",
		"plan": "application",
		"provider": null,
		"syslog_drain_url": null,
		"tags": [],
		"volume_mounts": []
	}
]
}`
