const { response } = require('express');

const Pool = require('pg').Pool;

const pool = new Pool({
  user: 'admin1',
  host: 'pmc1.ccyv4mlgftmr.us-east-1.rds.amazonaws.com',
  database: 'postgres',
  password: 'admin123',
  port: 5432,
});

// ----------------------------------------------- GET functions -------------------------------------------------------------------------------
const getBuildings = (request, response) => 
{
  pool.query('select * from public.building ORDER BY id ASC', (error, result) => {
    if(error)
    {
      response.status(400).json(error);
    }
    response.status(200).json(result.rows);
  })
}

const getClasses = (request, response) => 
{
  pool.query('select * from public.class ORDER BY id ASC', (error, result) => {
    if(error)
    {
      response.status(400).json(error);
    }
    response.status(200).json(result.rows);
  })
}

const getColleges = (request, response) => 
{
  pool.query('select * from public.college ORDER BY id ASC', (error, result) => {
    if(error)
    {
      response.status(400).json(error);
    }
    response.status(200).json(result.rows);
  })
}

const getCourses = (request, response) => 
{
  pool.query('select * from public.course ORDER BY id ASC', (error, result) => {
    if(error)
    {
      response.status(400).json(error);
    }
    response.status(200).json(result.rows);
  })
}

const getCourseSets = (request, response) => 
{
  pool.query('select * from public.course_set ORDER BY id ASC', (error, result) => {
    if(error)
    {
      response.status(400).json(error);
    }
    response.status(200).json(result.rows);
  })
}

const getCustomEvents = (request, response) => 
{
  pool.query('select * from public.custom_event ORDER BY id ASC', (error, result) => {
    if(error)
    {
      response.status(400).json(error);
    }
    response.status(200).json(result.rows);
  })
}

const getMajors = (request, response) => 
{
  pool.query('select * from public.major ORDER BY id ASC', (error, result) => {
    if(error)
    {
      response.status(400).json(error);
    }
    response.status(200).json(result.rows);
  })
}

const getOverallRatings = (request, response) => 
{
  pool.query('select * from public.over_all_rating ORDER BY id ASC', (error, result) => {
    if(error)
    {
      response.status(400).json(error);
    }
    response.status(200).json(result.rows);
  })
}

const getPrerequisites = (request, response) => 
{
  pool.query('select * from public.prerequisites ORDER BY id ASC', (error, result) => {
    if(error)
    {
      response.status(400).json(error);
    }
    response.status(200).json(result.rows);
  })
}

const getProfessors = (request, response) => 
{
  pool.query('select * from public.professor ORDER BY id ASC', (error, result) => {
    if(error)
    {
      response.status(400).json(error);
    }
    response.status(200).json(result.rows);
  })
}

const getReviews = (request, response) => 
{
  pool.query('select * from review ORDER BY id ASC', (error, result) => {
    if(error)
    {
      response.status(400).json(error);
    }
    response.status(200).json(result.rows);
  })
}

const getSchedules = (request, response) => 
{
  pool.query('select * from public.schedule ORDER BY id ASC', (error, result) => {
    if(error)
    {
      response.status(400).json(error);
    }
    response.status(200).json(result.rows);
  })
}

const getSemesters = (request, response) => 
{
  pool.query('select * from public.semester ORDER BY id ASC', (error, result) => {
    if(error)
    {
      response.status(400).json(error);
    }
    response.status(200).json(result.rows);
  })
}

const getTags = (request, response) => 
{
  pool.query('select * from tag ORDER BY id ASC', (error, result) => {
    if(error)
    {
      response.status(400).json(error);
    }
    response.status(200).json(result.rows);
  })
}

const getUsers = (request, response) => 
{
  pool.query('select * from public.user ORDER BY id ASC', (error, result) => {
    if(error)
    {
      response.status(400).json(error);
    }
    console.log(result);
    response.status(200).json(result.rows);
  })
}

const getUserCourseHistory = (request, response) => 
{
  pool.query('select * from public.user_course_history ORDER BY id ASC', (error, result) => {
    if(error)
    {
      response.status(400).json(error);
    }
    console.log(result);
    response.status(200).json(result.rows);
  })
}

const getUserVotedTags = (request, response) => 
{
  pool.query('select * from public.user_voted_tag ORDER BY id ASC', (error, result) => {
    if(error)
    {
      response.status(400).json(error);
    }
    response.status(200).json(result.rows);
  })
}

// ----------------------------------------------- DELETE functions -------------------------------------------------------------------------------
const deleteBuilding = (request, response) =>
{
  const id = parseInt(request.params.id)

  pool.query('DELETE FROM public.building WHERE id = $1', [id], (error, results) => {
    if (error) 
    {
      response.status(400).json(error);
    }
    response.status(201).send({"message" : "Delete data successfully"});
  })
}

const deleteClass = (request, response) =>
{
  const id = parseInt(request.params.id)

  pool.query('DELETE FROM public.class WHERE id = $1', [id], (error, results) => {
    if (error) 
    {
      response.status(400).json(error);
    }
    response.status(201).send({"message" : "Delete data successfully"});
  })
}

const deleteCollege = (request, response) =>
{
  const id = parseInt(request.params.id)

  pool.query('DELETE FROM public.college WHERE id = $1', [id], (error, results) => {
    if (error) 
    {
      response.status(400).json(error);
    }
    response.status(201).send({"message" : "Delete data successfully"});
  })
}

const deleteCourse = (request, response) =>
{
  const id = parseInt(request.params.id)

  pool.query('DELETE FROM public.course WHERE id = $1', [id], (error, results) => {
    if (error) 
    {
      response.status(400).json(error);
    }
    response.status(201).send({"message" : "Delete data successfully"});
  })
}

const deleteCourseSet = (request, response) =>
{
  const id = parseInt(request.params.id)

  pool.query('DELETE FROM public.course_set WHERE id = $1', [id], (error, results) => {
    if (error) 
    {
      response.status(400).json(error);
    }
    response.status(201).send({"message" : "Delete data successfully"});
  })
}

const deleteCustomEvent = (request, response) =>
{
  const id = parseInt(request.params.id)

  pool.query('DELETE FROM public.custom_event WHERE id = $1', [id], (error, results) => {
    if (error) 
    {
      response.status(400).json(error);
    }
    response.status(201).send({"message" : "Delete data successfully"});
  })
}


// ----------------------------------------------- Class functions -------------------------------------------------------------------------------
const updateClass = (request, response) => {
  const { id, created_at, deleted_at, is_deleted, semester, year, session, wait_list, offer_date, start_time, end_time, location, recommendation_score, 
  type, number, component, unit, seat_available, notes, instructors, course_title, course_catalog_name, course_id, rating } = request.body

  pool.query('UPDATE public.class SET created_at = $2, deleted_at = $3, is_deleted = $4, semester = $5, year = $6, session = $7, wait_list = $8, offer_date = $9, start_time = $10, end_time = $11, location = $12, recommendation_score = $13, type = $14, number = $15, component = $16, unit = $17, seat_available = $18, notes = $19, instructors = $20, course_title = $21, course_catalog_name = $22, course_id = $23, rating = $24 where id = $1',
              [id, created_at, deleted_at, is_deleted, semester, year, session, wait_list, offer_date, start_time, end_time, location, recommendation_score, 
                type, number, component, unit, seat_available, notes, instructors, course_title, course_catalog_name, course_id, rating], 
              (error, result) => {
                if (error) 
                {
                  response.status(400).json(error);
                }
                response.status(201).send({"message" : "Update data successfully"});
              }
  )
}



// ----------------------------------------------- College functions -------------------------------------------------------------------------------


const updateCollege = (request, response) => {
  const { id, created_at, deleted_at, is_deleted, name } = request.body

  pool.query('UPDATE public.college SET created_at = $2, deleted_at = $3, is_deleted = $4, name = $5 where id = $1',
              [id, created_at, deleted_at, is_deleted, name], 
              (error, result) => {
                if (error) 
                {
                  response.status(400).json(error);
                }
                response.status(201).send({"message" : "Update data successfully"});
              }
  )
}

// ----------------------------------------------- Course functions -------------------------------------------------------------------------------


// ----------------------------------------------- Subject functions -------------------------------------------------------------------------------
const getSubjects = (request, response) => 
{
  pool.query('select * from subject ORDER BY id ASC', (error, result) => {
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

  pool.query('DELETE FROM public.subject WHERE id = $1', [id], (error, results) => {
    if (error) 
    {
      response.status(400).json(error);
    }
    response.status(201).send({"message" : "Delete data successfully"});
  })
}

// ----------------------------------------------- Google Users functions -------------------------------------------------------------------------------


// ----------------------------------------------- Professor functions -------------------------------------------------------------------------------
const deleteProfessor = (request, response) =>
{
  const id = parseInt(request.params.id)

  pool.query('DELETE FROM public.professor WHERE id = $1', [id], (error, results) => {
    if (error) 
    {
      response.status(400).json(error);
    }
    response.status(201).send({"message" : "Delete data successfully"});
  })
}

const updateProfessor = (request, response) => {
  const { id, name, college_id, deleted_at} = request.body

  pool.query('UPDATE public.professor SET name = $2, deleted_at = $3, college_id = $4, deleted_at = $5 where id = $1',
              [id, name, deleted_at, college_id, deleted_at], 
              (error, result) => {
                if (error) 
                {
                  response.status(400).json(error);
                }
                response.status(201).send({"message" : "Update data successfully"});
              }
  )
}

// ----------------------------------------------- Review functions -------------------------------------------------------------------------------
const deleteReview = (request, response) =>
{
  const id = parseInt(request.params.id)

  pool.query('DELETE FROM public.review WHERE id = $1', [id], (error, results) => {
    if (error) 
    {
      response.status(400).json(error);
    }
    response.status(201).send({"message" : "Delete data successfully"});
  })
}

const updateReview = (request, response) => {
  const { id, created_at, deleted_at, is_deleted, rating, anonymous, recommended, pros, cons, comment, course_id, user_id, like_count, dislike_count } = request.body

  pool.query('UPDATE public.review SET created_at = $2, deleted_at = $3, is_deleted = $4, rating = $5, anonymous = $6, recommended = $7, pros = $8, cons = $9, comment = $10, course_id = $11, user_id = $12, like_count = $13, dislike_count = $14 where id = $1',
              [id, created_at, deleted_at, is_deleted, rating, anonymous, recommended, pros, cons, comment, course_id, user_id, like_count, dislike_count], 
              (error, result) => {
                if (error) 
                {
                  response.status(400).json(error);
                }
                response.status(201).send({"message" : "Update data successfully"});
              }
  )
}

// ----------------------------------------------- Semester functions -------------------------------------------------------------------------------


// ----------------------------------------------- Tag functions -------------------------------------------------------------------------------


// ----------------------------------------------- Overall Rating functions -------------------------------------------------------------------------------


// ----------------------------------------------- User Voted Tags functions -------------------------------------------------------------------------------


// ----------------------------------------------- User functions -------------------------------------------------------------------------------

const deleteUser = (request, response) =>
{
  const id = parseInt(request.params.id)

  pool.query('DELETE FROM public.user WHERE id = $1', [id], (error, results) => {
    if (error) 
    {
      response.status(400).json(error);
    }
    response.status(201).send({"message" : "Delete data successfully"});
  })
}

const updateUser = (request, response) => {
  const { id, created_at, deleted_at, is_deleted, first_name, last_name, password, email, college_id, academic_year, avatar, user_id, role } = request.body

  pool.query('UPDATE public.user SET created_at = $2, deleted_at = $3, is_deleted = $4, first_name = $5, last_name = $6, password = $7, email = $8, college_id = $9, academic_year = $10, avatar = $11, user_id = $12, role = $13 where id = $1',
              [id, created_at, deleted_at, is_deleted, first_name, last_name, password, email, college_id, academic_year, avatar, user_id, role], 
              (error, result) => {
                if (error) 
                {
                  response.status(400).json(error);
                }
                response.status(201).send({"message" : "Update data successfully"});
              }
  )
}

// ----------------------------------------------- Schedule functions -------------------------------------------------------------------------------
const deleteSchedule = (request, response) =>
{
  const id = parseInt(request.params.id)

  pool.query('DELETE FROM public.schedule WHERE id = $1', [id], (error, results) => {
    if (error) 
    {
      response.status(400).json(error);
    }
    response.status(201).send({"message" : "Delete data successfully"});
  })
}

const getScheduleByID = (request, response) => 
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

const updateSchedule = (request, response) => {
  const { id, user_id, class_id, semester_id, deleted_at, created_at, is_deleted } = request.body

  pool.query('UPDATE public.schedule SET user_id = $2, class_id = $3, semester_id = $4, deleted_at = $5, created_at = $6, is_deleted = $7 where id = $1',
              [id, user_id, class_id, semester_id, deleted_at, created_at, is_deleted], 
              (error, result) => {
                if (error) 
                {
                  response.status(400).json(error);
                }
                response.status(201).send({"message" : "Update data successfully"});
              }
  )
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

const deleteFeedback = (request, response) =>
{
  const id = parseInt(request.params.id)

  pool.query('DELETE FROM public.feedback WHERE id = $1', [id], (error, results) => {
    if (error) 
    {
      response.status(400).json(error);
    }
    response.status(201).send({"message" : "Delete data successfully"});
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

const updateFeedBack = (request, response) => {
  const { id, user_id, class_id, semester_id, rating, feedback } = request.body

  pool.query('UPDATE public.feedback SET user_id = $2, class_id = $3, semester_id = $4, rating = $5, feedback = $6 where id = $1',
              [id, user_id, class_id, semester_id, rating, feedback], 
              (error, result) => {
                if (error) 
                {
                  response.status(400).json(error);
                }
                response.status(201).send({"message" : "Update data successfully"});
              }
  )
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


module.exports = 
{
  getBuildings,
  getClasses,
  getColleges,
  getCourses,
  getCourseSets,
  getCustomEvents,
  getMajors,
  getOverallRatings,
  getPrerequisites,
  getProfessors,
  getReviews,
  getSchedules,
  getSemesters,
  getSubjects,
  getTags,
  getUsers,
  getUserCourseHistory,
  getUserVotedTags,
  getFeedbacks,

  getScheduleByID,
  addToSchedule,
  removeFromSchedule,
  updateSchedule,
  deleteSchedule,


  getFeedbackById,
  updateFeedBack,
  createFeedback,
  deleteFeedback,

  deleteClass,
  updateClass,

  deleteCollege,
  updateCollege,

  deleteCourse,


  deleteSubject,

  updateProfessor,
  deleteProfessor,

  updateReview,
  deleteReview,


  updateUser,
  deleteUser
}
