import express from 'express';
import bodyParser  from 'body-parser';
import { getSchedule, postSchedule} from './queries';

const app = express();
const PORT = process.env.PORT || 5000;

app.use(bodyParser.json());

app.use(
    bodyParser.urlencoded({
      extended: true,
    })
);

app.get('/', (request, response) => {
  response.json({ info: 'This is the schedule API.' })
})

app.get('/schedule/:id', getSchedule);
app.post('/schedule', postSchedule);

app.listen(PORT, () => {
  console.log(`App running on port ${PORT}.`)
})