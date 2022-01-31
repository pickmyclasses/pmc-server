import { Pool } from 'pg';

const pool = new Pool({
  user: 'admin1',
  host: 'pmc1.ccyv4mlgftmr.us-east-1.rds.amazonaws.com',
  database: 'postgres',
  password: 'admin123',
  port: 5432,
});

const getSchedule = (request, response) => 
{
  const id = parseInt(request.params.id)

  pool.query('SELECT * FROM schedule WHERE user_id = $1', [id], (error, results) => 
  {
    if (error) 
    {
      throw error;
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

export default 
{
  getSchedule,
  postSchedule
};