# Geeam

Geeam is a tool that helps administrators quickly find out which VMs are present in the Veean backup jobs.    
In some cases, especially when there is a large number of backup objects, it becomes difficult to find out 
which of those objects are included in Veeams backup jobs. This is even more emphasized when the VMs are 
constantly beeing added.      
This tool queries the Veeam servers inventory and displays the information on objects that are not included 
in any backup job present on the Veeam server.

## Configuration
* `-veeam-host` - the IP address / DNS name of the server that is running Veeam B&R
* `-veeam-username` - the username that can authenticate to Veeam B&R. This is usually, the admin user on that server.
* `-veeam-pass` - the password for Veeam server user

## Example
```shell
geeam -veeam-host veeam.server -veeam-username veeam.user -veeam-pass "VeeamPassw0rd"                   
VM Name           Backup  Host
Fr...             NO      10.x.x.x
ka...             NO      10.x.x.x  
AM...             NO      10.x.x.x 
UB...             NO      10.x.x.x 
AM...             NO      10.x.x.x  
GD...             NO      10.x.x.x 
UB...             NO      10.x.x.x  
AM...             NO      10.x.x.x 
ds...             NO      10.x.x.x 
```
