const express = require('express');
const bodyParser = require('body-parser');
const app = express();
const db = require('./queries');
const PORT = process.env.PORT || 5000;

// parse the body into json
app.use(bodyParser.json());
app.use(
  bodyParser.urlencoded({
    extended: true,
  })
);

app.listen(PORT, () => {
  console.log(`App running on port ${PORT}.`)
});

app.get('/', (request, response) => {
  response.json({ info: 'schedule API' })
});

app.get('/schedule', db.getSchedule);
app.post('/schedule', db.postSchedule);

