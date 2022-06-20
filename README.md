<h3>Backend for @TheHubAUBG's SPA</h3>
<hr/>
<p><strong>Information and resources: </strong></p>
<p><a href="https://github.com/asynchroza/Hub-Website-Backend/blob/main/tasks.txt">Timeline of added features</a></p>
<hr/>
<h3>Project structure:</h3>
<p><strong>Admins (used for authorization): </p></strong>
<p>Requests: </p> 
<li> Post request on /api/login (login) - accepts username and password as body, and returns bearer token on success </li>
<hr/>
<p><strong>Members: (used for managing club members)</p></strong>
<p>Requests:</p>
<li> Post request on /api/member (create member) - accepts bearer token as a header, <br/> firstname, lastname, department, position, sociallink, profilepicture as body </li>
<li> Get request on /api/members (get all members) - accepts bearer token as a header </li>
