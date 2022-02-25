const express = require('express');
const bodyParser = require('body-parser');
const cors = require('cors');
const app = express();
const db = require('./functions');
const PORT = process.env.PORT || 5000;

app.use(cors());

// parse the body into json
app.use(bodyParser.json());
app.use(
  bodyParser.urlencoded({
    extended: true,
  })
);

app.listen(PORT, () => {
  console.log(`App running on port ${PORT}.`);
});

app.get('/', (request, response) => {
  response.json({ info: 'schedule API' });
});

// class API
app.get('/class', cors(), db.getClasses);

// college API
app.get('/college', cors(), db.getColleges);

// course API
app.get('/course', cors(), db.getCourses);

// course API
app.get('/subject', cors(), db.getSubjects);


// schedule API
app.get('/schedule/:user_id/:semester_id', cors(), db.getSchedule);
app.post('/schedule/add', db.addToSchedule);
app.post('/schedule/remove', db.removeFromSchedule);

// feedback API
app.get('/feedback', db.getFeedbacks);
app.get('/feedback/:id', cors(), db.getFeedbackById);
app.post('/feedback', db.createFeedback);
app.post('/feedback/remove/:id', cors(), db.deleteFeedback);
