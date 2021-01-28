Feature: Validate API responses
  DOCUMENTO_PROGRAMA_CRUD
  probe JSON reponses


Scenario Outline: To probe route code response  /documento_programa
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>"      

  Examples: 
  |method|route                 |bodyreq               |codres       |
  |GET   |/v1/documento_programa|./files/req/Vacio.json|200 OK       |
  |GET   |/v1/documento_program |./files/req/Vacio.json|404 Not Found|
  |POST  |/v1/documento_program |./files/req/Vacio.json|404 Not Found|
  |PUT   |/v1/documento_program |./files/req/Vacio.json|404 Not Found|
  |DELETE|/v1/documento_program |./files/req/Vacio.json|404 Not Found|


Scenario Outline: To probe response route /documento_programa       
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>"      
  And the response should match json "<bodyres>"

  Examples: 
  |method|route                 |bodyreq               |codres         |bodyres                |
  |GET   |/v1/documento_programa|./files/req/Vacio.json|200 OK         |./files/res1/Vok1.json |
  |POST  |/v1/documento_programa|./files/req/Vacio.json|400 Bad Request|./files/res1/Ierr1.json|
 