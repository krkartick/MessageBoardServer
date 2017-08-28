Environment setup and steps to test this project:
	1. Install Docker on the test machine 
	
	2. Pull the docker image from following path:
	https://hub.docker.com/r/krkartick/mbserverv1.project/
	Or 
	https://hub.docker.com/r/krkartick/mbserverv2.project/
	
	3. Install MySQL on any machine. Note the IP Address of the MySQL Database
	4. Create the schema and database using the script <MySQL_DB_Structure.sql>
	5. Update the MySQL DB IP address in <RunDockerApp.sh>
	6. Run the docker image using <RunDockerApp.sh>
	7. Now, 
		the <mbserverv1> runs on port 9000 Or
			<mbserverv2> runs on port 9900
	8. From any browser REST can be tested with "IPAddress" of the Host machine 
		URL Syntax: 
		"http:/<IPAddress>:9000/Topics/<topicname>"
		
		Examples:
		HTTP GET 	http://192.168.56.5:9000/Topics/Cat		-> Retrieves Posts on Cat 
		HTTP POST 	http://192.168.56.5:9000/Topics/Tiger 	-> Creates a Topic Tiger 
		HTTP POST 	http://192.168.56.5:9000/Topics/Tiger 
			 BODY 	"Tiger belongs to Cat family" 			-> Posts the message on Topic Tiger
		HTTP DELETE	http://192.168.56.5:9000/Topics/Tiger 	-> Deletes the Topic Tiger 

		Recommened REST testing tools chrome extensions:
			Restlet Client 
			RESTful Stress - For performance Testing 