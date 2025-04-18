# 2nd-Midterm-OS
1. Dockerfile 
To avoid the 'wget not found' error or similar issues, I decided to install it using the following command:

-> apt-get install wget

After getting the following error: 

-> Do you want to continue? [Y/n] Abort.

I added the -y flag to confirm the installation. 
-> apt-get -y install wget. 


Now that we have the .zip file  its time to unzip it.

-> unzip /tmp/godot.zip \ . 

Rename it 
-> mv Godot_v4.4-stable_linux.arm64 /usr/local/bin/godot \

and remove the zip. 

-> && rm /tmp/godot.zip

When trying to run my image, got the following error. 
rosetta error: failed to open elf at /lib64/ld-linux-x86-64.so.

This error occurred because I’m working on a computer with ARM architecture, so I decided to create two images: one for ARM and another for x86.

I changed the download link for Godot to one that matched my architecture and rebuilt the image. Now that I have two Dockerfiles, when running the docker build command, I need to specify which file to use. To do that, I used the following command: 

-> docker build -f Dockerfile.ARM . 

Source: https://docs.docker.com/reference/cli/docker/buildx/build/


However, I encountered the following error:

-> ERROR: Unable to load fontconfig, system font support is disabled.

To fix this, I added fontconfig to the installation line. Now, the container runs correctly. 

 
2. Go Server and React 

First I created a Go server that could access a database following the official Go documentation:
https://go.dev/doc/tutorial/database-access

Decided to use a mysql database, modeling a mini figures store. Wrote a query to obtain the top ten most sold products depending on the year. 

SELECT productName, YEAR(o.orderDate) as year,COUNT(od.quantityOrdered) as quantity
FROM products p, orderdetails od, orders o
WHERE p.productCode = od.productCode AND
      od.orderNumber = o.orderNumber
GROUP BY p.productCode, year
HAVING (year = '2003')
ORDER BY year,quantity DESC
LIMIT 10

Then, I transformed my server into an API using the following documentation:
https://go.dev/doc/tutorial/web-service-gin

Created the react front using the following sources: 
- https://www.youtube.com/watch?v=ZiEDNz6GbQM&t=42s


Sources: 
- https://docs.docker.com/get-started/docker-concepts/building-images/build-tag-and-publish-an-image/
- https://docs.docker.com/reference/dockerfile/
- https://docs.docker.com/build/building/best-practices/
- https://docs.docker.com/engine/containers/run/#image-references
- https://docs.docker.com/engine/containers/run/
- https://docs.docker.com/get-started/docker-concepts/building-images/writing-a-dockerfile/
- https://docs.docker.com/get-started/docker-overview/
- https://docs.docker.com/compose/intro/compose-application-model/
- https://docs.docker.com/compose/gettingstarted/
- https://docs.godotengine.org/en/stable/about/introduction.html
- https://linux.die.net/man/1/wget
- https://www.docker.com/blog/how-to-dockerize-react-app/


