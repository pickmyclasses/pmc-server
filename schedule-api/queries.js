const Pool = require('pg').Pool;

const pool = new Pool({
  user: 'admin1',
  host: 'pmc1.ccyv4mlgftmr.us-east-1.rds.amazonaws.com',
  database: 'postgres',
  password: 'admin123',
  port: 5432,
});

const getSchedule = (request, response) => 
{
  const { user_id, semester_id } = request.body;

  pool.query('SELECT * FROM class INNER JOIN schedule ON class.id = schedule.class_id WHERE schedule.user_id = $1 AND schedule.semester_id = $2', 
              [user_id, semester_id], (error, results) => 
  {
    if (error) 
    {
      response.status(500).json(error);
    }
    response.status(200).json(results.rows);
  })
}

const postSchedule = (request, response) => 
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

module.exports = {
  getSchedule,
  postSchedule,
}