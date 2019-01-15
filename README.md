# go-RestFull-ORM

This example is about RestFull Web service in Go Lang interacting with DataBase (Postgresql):
74
weservcies:
	/api/findall- GET 
	/api/create - POST 
	/api/update - PUT
	/api//delete/{id} - DELETE 
	/api/exit - GET 
	
	
Example also includes LoggingHandler which will Log Each and every Service call in a proper fashion.


How to Start:

	Helm chart contains Webservice in go-lang which will interact with DB (postgresql)
		
	1:Using Helm Chart inside K8
	
		Download the helm chart from my github.
		
				helm package /path/to/my-helm-chart
				helm install /output/of/package.tgz
		
		Verify Installion:
		
				kubectl get pods 
				kubectl get services
		
		Get Port of the service mapping and verify with : http://machineIp:servicePort:/api/exit 
