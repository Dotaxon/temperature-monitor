## THIS README was written for my own RPI-Heizung raspberry pi zero
**therefore it will not apply to you**


go api is under /etc/GO-API started via cornjob 
new go api has to be build locally and then moved to /etc/GO-API make sure that the parameters in main.go are right

the website is under /var/www/html it can be build offsite with e.g. npm run build and then copied to /var/www/html
the configuration is for apache is in /etc/apache2
the /etc/apache2/sites-enabled folder shows the enabled configs
to change a config first make the changes in /etc/apache2/sites-available/myown-ssl.conf
there are commands that check the syntax of the config
 
then with sudo rights:
a2dissite -p myown-ssl.conf #disable config, the -p purges any remaing traces of the config or something like that
a2ensite myown-ssl.conf #enable config
systemctl reload apache2
systemclt status apache2 #check if everything works

### to renew certificate
1. `sudo certbot certonly --manual -d "*.<Domain>"`
2. save certs to xca
3. transfer certs to rpi-heizung (and other places where it is used)
4. save certs to /etc/ssl/...
5. execute `update-ca-certificates`
6. make sure main.go also uses the right certs and maybe recomiple it with `go build`
7. wipe the cert in every place you copied it to (except the place where certbot created it)
