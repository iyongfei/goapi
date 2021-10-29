#!/bin/bash

list="rules/bypass rules/chat rules/cmi rules/denial_of_service rules/dhcp
 rules/dnp3 rules/dns rules/dos rules/flood rules/ftp rules/fuzz rules/game
 rules/http rules/hunting rules/icmp rules/iis rules/imap rules/info rules/js
 rules/kerberos rules/microsoft rules/netbios rules/nfs rules/noop rules/ntp
 rules/oracle rules/overflow rules/php rules/policy rules/portscan rules/shellcode
 rules/sip  rules/smb  rules/smtp  rules/snmp rules/sql rules/ssh rules/ssl
 rules/telnet rules/tftp rules/tls rules/trojan rules/vbs rules/web rules/worm
 rules/xss"
for i in $list;
do
  ./xml -v $i
done
