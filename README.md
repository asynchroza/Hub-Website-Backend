# Backend for @TheHubAUBG's SPA

<a href="https://github.com/asynchroza/Hub-Website-Backend/blob/main/go.mod"> Dependencies </a>
---
<p><strong>Information and resources: </strong></p>
<p><a href="https://github.com/asynchroza/Hub-Website-Backend/blob/main/tasks.txt">Timeline of added features</a></p>

---
### Endpoint structure:

<p><strong>NB:</strong> bearer_token is not passed in Authorization header but in a custom basic one - "BEARER_TOKEN" </p>
<p><strong>Admins (used for authorization): </p></strong>
<p>Requests: </p> 
<li> Post request on /api/login (login) - accepts username and password as body, and returns bearer token on success </li>
<li> Post request on /api/validate (BEARER_TOKEN validation) - accepts header with BEARER_TOKEN </li>
<hr/>
<p><strong>Members (used for managing club members):</p></strong>
<p>Requests:</p>
<li> Post request on /api/member (create member) - accepts bearer token as a header, member body as form data </li>
<li> Get request on /api/members (get all members) - empty request, returns all members </li>
<li> Get request on /api/member/:key (get single member) - accepts parameter key (memberid) (pass as ../member/1 not ../?key=1)</li>
<li> Put request on /api/member/:key (change info of member) - accepts parameter key (memberid), event body as form data and bearer token</li>
<li> Delete request on /api/member/:key (delete member) - accepts parameter key (memberid) and bearer token as a header</li>
<br>
<p><em><strong>NB:</strong> GET requests are not subject to authorization</em></p>
<hr/>
<p><strong>Events (used for managing displayed events): </p></strong>
<p><em><a href="http://www.timestamp-converter.com/">Use this website to convert dates to ISO format</a></em></p>
<li> Post request on /api/event (create event) - accepts bearer token as a header and event body as form data</li>
<li> Get request on /api/event/:key (get event) - accepts key parameter and returns indexed event (eventid) </li>
<li> Get request on /api/events (get all events) - no parameters needed, returns all events </li>
<li> Put request on /api/event/:key (edit event) - accepts bearer_token, event body as form data </li>
<hr/>
<p><strong>Articles (used for managing displayed articles): </p></strong>
<p> TBA </p>
<hr/>
<p><strong>Jobs (used for managing displayed job opportunities): </p></strong>
<p> TBA </p>
