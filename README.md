<h3>Backend for @TheHubAUBG's SPA</h3>
<p><em>Go, MongoDB</em></p>
<a href="https://github.com/asynchroza/Hub-Website-Backend/blob/main/go.mod"> Dependencies </a>
<hr/>
<p><strong>Information and resources: </strong></p>
<p><a href="https://github.com/asynchroza/Hub-Website-Backend/blob/main/tasks.txt">Timeline of added features</a></p>
<hr/>
<h3>Endpoint structure:</h3>
<p><strong>Admins (used for authorization): </p></strong>
<p>Requests: </p> 
<li> Post request on /api/login (login) - accepts username and password as body, and returns bearer token on success </li>
<hr/>
<p><strong>Members (used for managing club members):</p></strong>
<p>Requests:</p>
<li> Post request on /api/member (create member) - accepts bearer token as a header, <br/> firstname, lastname, department, position, sociallink, profilepicture as body </li>
<li> Get request on /api/members (get all members) - accepts bearer token as a header (might remove the validation for get requests) </li>
<li> Get request on /api/member (get single member) - TBA </li>
<li> Put request on /api/member (change info of member) - TBA </li>
<li> Delete request on /api/member (delete member) - TBA </li>
<hr/>
<p><strong>Events (used for managing displayed events): </p></strong>
<p> TBA </p>
<hr/>
<p><strong>Articles (used for managing displayed articles): </p></strong>
<p> TBA </p>
<hr/>
<p><strong>Jobs (used for managing displayed job opportunities): </p></strong>
<p> TBA </p>
