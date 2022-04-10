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
  response.json({ info: 'Admin Data Table API' });
});

// GET APIs
app.get('/building', cors(), db.getBuildings);
app.get('/class', cors(), db.getClasses);
app.get('/college', cors(), db.getColleges);
app.get('/course', cors(), db.getCourses);
app.get('/course_set', cors(), db.getCourseSets);
app.get('/custom_event', cors(), db.getCustomEvents);
app.get('/major', cors(), db.getMajors);
app.get('/over_all_rating', cors(), db.getOverallRatings);
app.get('/prerequisites', cors(), db.getPrerequisites);
app.get('/professor', cors(), db.getProfessors);
app.get('/review', cors(), db.getReviews);
app.get('/schedule', cors(), db.getSchedules);
app.get('/semester', cors(), db.getSemesters);
app.get('/subject', cors(), db.getSubjects);
app.get('/tag', cors(), db.getTags);
app.get('/user', cors(), db.getUsers);
app.get('/user_course_history', cors(), db.getUserCourseHistory);
app.get('/user_voted_tag', cors(), db.getUserVotedTags);

// DELETE APIs
app.get('/building/delete/:id', cors(), db.deleteBuilding);
app.get('/class/delete/:id', cors(), db.deleteClass);
app.get('/college/delete/:id', cors(), db.deleteCollege);


// class API
// app.post('/class/update', cors(), db.updateClass);
// app.get('/class/delete/:id', cors(), db.deleteClass);

// college API

// app.post('/college/update', cors(), db.updateCollege);
// app.get('/college/delete/:id', cors(), db.deleteCollege);

// course API


// subject API


// professor API

// app.post('/professor/update', cors(), db.updateProfessor);
// app.get('/professor/delete/:id', cors(), db.deleteProfessor);

// review API

// app.post('/review/update', cors(), db.updateReview);
// app.get('/review/delete/:id', cors(), db.deleteReview);

// semester API


// user API

// app.post('/user/update', cors(), db.updateUser);
// app.get('/user/delete/:id', cors(), db.deleteUser);

// schedule API

// app.get('/schedule/:user_id/:semester_id', cors(), db.getSchedule);
// app.post('/schedule/add', db.addToSchedule);
// app.post('/schedule/remove', db.removeFromSchedule);
// app.post('/schedule/update', cors(), db.updateSchedule);
// app.get('/schedule/delete/:id', cors(), db.deleteSchedule);

// feedback API
// app.get('/feedback', db.getFeedbacks);
// app.get('/feedback/:id', cors(), db.getFeedbackById);
// app.post('/feedback/update', cors(), db.updateFeedBack);
// app.post('/feedback', db.createFeedback);
// app.get('/feedback/delete/:id', cors(), db.deleteFeedback);
