const http = require('http');

const connection = {
  hostname: 'localhost',
  port: '8000'
};

// response sender

const makeRquest = (options, responseJSON = null) => new Promise((res, rej) => {
  const { method } = options;
  const responseStringified = JSON.stringify(responseJSON);
  const resData = {};
  const responseOptions = { ...connection, ...options };
  if (method === 'POST')
    responseOptions.headers = {
      "Content-Type": "application/json",
      "Content-Length": Buffer.byteLength(responseStringified)
    };
  const servRespose = http.request(responseOptions, response => {
    response
      .setEncoding('utf-8')
      .on('data', data =>  (
        /^200|201|400|404|500$/.test(resData.status) ?
          resData.body = JSON.parse(data) :
          resData.body = data
      ))
      .on('close', () => (
        resData.status === 200 || 
        resData.status === 201 ?
          res(resData) : rej(resData)
      ))
  });
  servRespose
    .on('response', r => (resData.status = r.statusCode))
    .on('error', e => rej(e));
  if (method === 'POST')
    servRespose.end(responseStringified);
  else
    servRespose.end();
});


const toIncreaseVolume = (id1, id2)  => makeRquest(
  { method: 'POST', path: '/toIncreaseVolume' },
  [ id1, id2 ]
);

const findByIdMachine = id => makeRquest(
  { method: 'POST', path: '/findByIdMachine' }, 
  { id }
);

const findByIdDisk = id => makeRquest(
  { method: 'POST', path: '/findByIdDisk' },
  { id }
);



const getMachines = () => makeRquest({ method: 'GET', path: '/getMachines' });
const getDisks = () => makeRquest({ method: 'GET', path: '/getDisks' });





module.exports = {
  toIncreaseVolume,
  getMachines,
  getDisks,
  findByIdMachine,
  findByIdDisk
}