package main

import (
	"gopkg.in/ldap.v3"
	"log"
)

func main() {

	//gitlab_rails['ldap_servers'] = {
	//	'main' => {
	//		'label' => 'RITS AD',
	//		'host' =>  '10.20.8.15',
	//		'port' => 389,
	//	    'uid' => 'sAMAccountName',
	//		'allow_username_or_email_login' => true,
	//		'encryption' => 'plain',
	//		'verify_certificates' => true,
	//		'bind_dn' => 'CN=Service DevOps,OU=Service Accounts,OU=RITS User Accounts,DC=rits,DC=local',
	//		'password' => 'G0jhft671H',
	//		'active_directory' => true,
	//		'block_auto_created_users' => false,
	//		'base' => 'OU=RITS User Accounts,DC=rits,DC=local',
	//		'group_base' => 'OU=RITS Security Groups,OU=RITS User Accounts,DC=rits,DC=local',
	//		'admin_group' => 'CN=rits_devops_Administrators,OU=RITS Security Groups,OU=RITS User Accounts,DC=rits,DC=local'
	//   }
	//}

	//mapped 10.20.8.15:389 to 127.0.0.1:10389 locally because of firewall constraints
	host := "127.0.0.1"
	port := "10389"

	user := "svc_devops"
	domain := "@rits.local"
	password := "G0jhft671H"

	// TLS, for testing purposes disable certificate verification,
	// check https://golang.org/pkg/crypto/tls/#Config for further information.

	//tlsConfig := &tls.Config{InsecureSkipVerify: true}
	//l, err := ldap.DialTLS("tcp", "ldap.example.com:636", tlsConfig)

	// No TLS, not recommended
	l, err := ldap.Dial("tcp", host+":"+port)

	if err != nil {
		log.Println(err)
	} else {
		//Now you should have an active connection to your LDAP server.
		// Using this connection you have to execute a bind:
		//err = l.Bind("user@test.com", "password")
		err = l.Bind(user+domain, password)
		if err != nil {
			// error in ldap bind
			log.Println(err)
		}
		// successful bind
	}

}
