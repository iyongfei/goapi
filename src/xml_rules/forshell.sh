#!/bin/bash

list="rules/ftp rules/telnet rules/tftp rules/imap rules/smtp
 rules/dhcp rules/dnp3 rules/dns rules/http rules/icmp rules/kerberos rules/netbios
 rules/nfs rules/ntp rules/smb rules/ssl rules/snmp rules/ssh rules/tls
 rules/xss rules/shellcode rules/sql rules/dos rules/bypass rules/chat
 rules/cmi rules/flood rules/fuzz rules/game rules/hunting rules/info
 rules/sip  rules/trojan  rules/worm  rules/portscan rules/iis rules/js rules/noop
 rules/oracle rules/overflow rules/php rules/policy rules/vbs rules/web rules/microsoft
 rules/denial_of_service"
for i in $list;
do
  ./rule -v $i
done
