const { response } = require('express');

const Pool = require('pg').Pool;

const pool = new Pool({
  user: 'admin1',
  host: 'pmc1.ccyv4mlgftmr.us-east-1.rds.amazonaws.com',
  database: 'postgres',
  password: 'admin123',
  port: 5432,
});

// ----------------------------------------------- Class functions -------------------------------------------------------------------------------
const getClasses = (request, response) => 
{
  pool.query('select * from class', (error, result) => {
    if(error)
    {
      response.status(400).json(error);
    }
    response.status(200).json(result.rows);
  })
}

const deleteClass = (request, response) =>
{
  const id = parseInt(request.params.id)

  pool.query('DELETE FROM class WHERE id = $1', [id], (error, results) => {
    if (error) 
    {
      response.status(400).json(error);
    }
    response.status(200).json(results.rows)
  })
}

// ----------------------------------------------- College functions -------------------------------------------------------------------------------
const getColleges = (request, response) => 
{
  pool.query('select * from college', (error, result) => {
    if(error)
    {
      response.status(400).json(error);
    }
    response.status(200).json(result.rows);
  })
}

const deleteCollege = (request, response) =>
{
  const id = parseInt(request.params.id)

  pool.query('DELETE FROM college WHERE id = $1', [id], (error, results) => {
    if (error) 
    {
      response.status(400).json(error);
    }
    response.status(200).json(results.rows)
  })
}

// ----------------------------------------------- Course functions -------------------------------------------------------------------------------
const getCourses = (request, response) => 
{
  pool.query('select * from course', (error, result) => {
    if(error)
    {
      response.status(400).json(error);
    }
    response.status(200).json(result.rows);
  })
}

const deleteCourse = (request, response) =>
{
  const id = parseInt(request.params.id)

  pool.query('DELETE FROM course WHERE id = $1', [id], (error, results) => {
    if (error) 
    {
      response.status(400).json(error);
    }
    response.status(200).json(results.rows)
  })
}

// ----------------------------------------------- Subject functions -------------------------------------------------------------------------------
const getSubjects = (request, response) => 
{
  pool.query('select * from subject', (error, result) => {
    if(error)
    {
      response.status(400).json(error);
    }
    response.status(200).json(result.rows);
  })
}

const deleteSubject = (request, response) =>
{
  const id = parseInt(request.params.id)

  pool.query('DELETE FROM subject WHERE id = $1', [id], (error, results) => {
    if (error) 
    {
      response.status(400).json(error);
    }
    response.status(200).json(results.rows)
  })
}

// ----------------------------------------------- Google Users functions -------------------------------------------------------------------------------
const getGoogleUsers = (request, response) => 
{
  pool.query('select * from google_user', (error, result) => {
    if(error)
    {
      response.status(400).json(error);
    }
    response.status(200).json(result.rows);
  })
}

// ----------------------------------------------- Review functions -------------------------------------------------------------------------------
const getReviews = (request, response) => 
{
  pool.query('select * from review', (error, result) => {
    if(error)
    {
      response.status(400).json(error);
    }
    response.status(200).json(result.rows);
  })
}

// ----------------------------------------------- Review functions -------------------------------------------------------------------------------
const getSemesters = (request, response) => 
{
  pool.query('select * from semesters', (error, result) => {
    if(error)
    {
      response.status(400).json(error);
    }
    response.status(200).json(result.rows);
  })
}

// ----------------------------------------------- User functions -------------------------------------------------------------------------------
const getUsers = (request, response) => 
{
  pool.query('select * from users', (error, result) => {
    if(error)
    {
      response.status(400).json(error);
    }
    response.status(200).json(result.rows);
  })
}

// ----------------------------------------------- Schedule functions -------------------------------------------------------------------------------
const getSchedule = (request, response) => 
{
  const user_id = parseInt(request.params.user_id);
  const semester_id = parseInt(request.params.semester_id);

  pool.query('SELECT class_id FROM schedule WHERE schedule.user_id = $1 AND schedule.semester_id = $2', 
              [user_id, semester_id], (error, results) => 
  {
    if (error) 
    {
      response.status(500).json(error);
    }
    response.status(200).json(results.rows);
  })
}

const addToSchedule = (request, response) => 
{
  const { user_id, class_id, semester_id } = request.body;

  pool.query('INSERT INTO schedule (user_id, class_id, semester_id) VALUES ($1, $2, $3)', [user_id, class_id, semester_id], (error, results) => 
  {
    if (error) 
    {
      throw error;
    }
    response.status(201).send(`schedule added for user with ID: ${request.body.user_id}`);
  });
}

const removeFromSchedule = (request, response) => 
{
  const { user_id, class_id, semester_id } = request.body;

  pool.query('DELETE FROM schedule WHERE user_id = $1 AND class_id = $2 AND semester_id = $3;', [user_id, class_id, semester_id], (error, results) => 
  {
    if (error) 
    {
      response.status(400).json(error);
    }
    response.status(201).send(`The class has been removed from the schedule`);
  });  
}


// ----------------------------------------------- Feedback functions -------------------------------------------------------------------------------
const getFeedbacks = (request, response) => {
  pool.query('SELECT * FROM feedback ORDER BY id ASC', (error, results) => {
    if (error) {
      response.status(400).json(error);
    }
    response.status(200).json(results.rows);
  })
}

const getFeedbackById = (request, response) => {
  const id = parseInt(request.params.id)

  pool.query('SELECT * FROM feedback WHERE id = $1', [id], (error, results) => {
    if (error) 
    {
      response.status(400).json(error);
    }
    response.status(200).json(results.rows)
  })
}

const createFeedback = (request, response) => {
  const { user_id, class_id, semester_id, rating, feedback } = request.body

  pool.query('INSERT INTO feedback (user_id, class_id, semester_id, rating, feedback) VALUES ($1, $2, $3, $4, $5)', 
                                   [user_id, class_id, semester_id, rating, feedback], (error, results) => 
  {
    if (error) 
    {
      response.status(400).json(error);
    }
    response.status(201).send(`Feedback added successfully`)
  })
}

const deleteFeedback = (request, response) =>
{
  const id = parseInt(request.params.id)

  pool.query('DELETE FROM feedback WHERE id = $1', [id], (error, results) => {
    if (error) 
    {
      response.status(400).json(error);
    }
    response.status(200).json(results.rows)
  })
}


module.exports = 
{
  getSchedule,
  addToSchedule,
  removeFromSchedule,

  getFeedbacks,
  getFeedbackById,
  createFeedback,
  deleteFeedback,

  getClasses,
  deleteClass,

  getColleges,
  deleteCollege,

  getCourses,
  deleteCourse,

  getSubjects,
  deleteSubject,

  getGoogleUsers,
  getReviews,
  getSemesters,
  getUsers
}
