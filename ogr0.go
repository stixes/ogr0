package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"github.com/Sirupsen/logrus"
	"golang.org/x/crypto/ssh"
)

const appName = "ogr0"
const host_key_path = "host_key"
const bind_addr = ":2222"
const server_version_string = "SSH-2.0-OpenSSH_8.4p1 Ubuntu-6ubuntu2.1"

var errAuthenticationFailed = errors.New(":)")

func logParameters(conn ssh.ConnMetadata) logrus.Fields {
	return logrus.Fields{
		"user":                conn.User(),
		"client_version":      string(conn.ClientVersion()),
		"remote_addr_address": string(conn.RemoteAddr().String()),
	}
}

func authenticatePassword(conn ssh.ConnMetadata, password []byte) (*ssh.Permissions, error) {
	logrus.WithFields(logParameters(conn)).Info(fmt.Sprintf("Request with password: %s ", password))
	return nil, errAuthenticationFailed
}

func authenticateKey(conn ssh.ConnMetadata, key ssh.PublicKey) (*ssh.Permissions, error) {
	logrus.WithFields(logParameters(conn)).Info(fmt.Sprintf("Request with keytype: %s ", key.Type()))
	return nil, errAuthenticationFailed
}

func main() {
	config := ssh.ServerConfig{
		PasswordCallback:  authenticatePassword,
		PublicKeyCallback: authenticateKey,
		ServerVersion: server_version_string,
	}
	hostPrivateKey, err := ioutil.ReadFile(host_key_path)
	if err != nil {
		logrus.Panic(err)
	}
	hostPrivateKeySigner, err := ssh.ParsePrivateKey(hostPrivateKey)
	if err != nil {
		logrus.Panic(err)
	}
	config.AddHostKey(hostPrivateKeySigner)
	socket, err := net.Listen("tcp", bind_addr)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := socket.Accept()
		if err != nil {
			logrus.Panic(err)
		}
		_, _, _, err = ssh.NewServerConn(conn, &config)
		if err == nil {
			logrus.Panic("Successful login? why!?")
		}
	}
}

