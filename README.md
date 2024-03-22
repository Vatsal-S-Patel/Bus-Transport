# Bus Transport Project

## Statement:
Users can utilize our app to search for available BRTS (Bus Rapid Transit System) stations near them. They will receive information on available routes and incoming buses. Additionally, they can opt to receive live bus locations on a map, which is highly beneficial for users. If direct buses are not available between station X and Y, the app will display alternative routes via buses, including their route names and junction station names. 

Admins have the capability to view and modify all routes, schedule buses, and assign drivers to buses, ensuring efficient management of the system. Therefore, this application proves to be invaluable for city commuters.

## Work Done:
We have leveraged the `net/http` package in Golang for server implementation. Additionally, we have incorporated Socket.IO for real-time communication within the Golang environment. Complex joins in PostgreSQL database have been utilized for managing routes and indirect buses efficiently. Views have been employed to optimize performance and avoid repetitive complex joins.