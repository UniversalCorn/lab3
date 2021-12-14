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

// client requests

// const addForum = (name, topic) => makeRquest(
//   { method: 'POST', path: '/forums' },
//   { name, topic }
// );

// const addUser = (name, interests) => makeRquest(
//   { method: 'POST', path: '/users' },
//   { name, interests }
// );

// const getMachine = name => makeRquest(
//   { method: 'POST', path: '/machine' },
//   { name }
// );

// const getDisk = name => makeRquest(
//   { method: 'POST', path: '/disk' },
//   { name }
// );

const getMachines = () => makeRquest({ method: 'GET', path: '/machines' });
const getDisks = () => makeRquest({ method: 'GET', path: '/discks' });

module.exports = {
//   addForum,
//   getMachine,
  getMachines,
 // getDisk,
  getDisks
}