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
app.post('/class/update', cors(), db.updateClass);
app.get('/class/delete/:id', cors(), db.deleteClass);

// college API
app.get('/college', cors(), db.getColleges);
app.post('/college/update', cors(), db.updateCollege);
app.get('/college/delete/:id', cors(), db.deleteCollege);

// course API
app.get('/course', cors(), db.getCourses);

// subject API
app.get('/subject', cors(), db.getSubjects);

// professor API
app.get('/professor', cors(), db.getProfessors);
app.post('/professor/update', cors(), db.updateProfessor);
app.get('/professor/delete/:id', cors(), db.deleteProfessor);

// google_user API
app.get('/google_user', cors(), db.getGoogleUsers);

// review API
app.get('/review', cors(), db.getReviews);
app.post('/review/update', cors(), db.updateReview);
app.get('/review/delete/:id', cors(), db.deleteReview);

// semester API
app.get('/semester', cors(), db.getSemesters);

// user API
app.get('/user', cors(), db.getUsers);
app.post('/user/update', cors(), db.updateUser);
app.get('/user/delete/:id', cors(), db.deleteUser);

// schedule API
app.get('/schedule', cors(), db.getSchedules);
app.get('/schedule/:user_id/:semester_id', cors(), db.getSchedule);
app.post('/schedule/add', db.addToSchedule);
app.post('/schedule/remove', db.removeFromSchedule);
app.post('/schedule/update', cors(), db.updateSchedule);
app.get('/schedule/delete/:id', cors(), db.deleteSchedule);

// feedback API
app.get('/feedback', db.getFeedbacks);
app.get('/feedback/:id', cors(), db.getFeedbackById);
app.post('/feedback/update', cors(), db.updateFeedBack);
app.post('/feedback', db.createFeedback);
app.get('/feedback/delete/:id', cors(), db.deleteFeedback);
