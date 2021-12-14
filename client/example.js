const machineClient = require('./machinesClinet');


const testData = [
    ['getMachines', 'Getting all machines', []],
    ['getDisks', 'Getting all disks', []],
    ['findByIdMachine', 'Find Machine by ID withour ID', []],
    ['findByIdMachine', 'Find Machine by ID: 1', [1]],
    ['findByIdMachine', 'Find Machine by ID: 2', [2]],
    ['findByIdMachine', 'Find Machine by ID: 3', [3]],
    ['findByIdDisk', 'Find Disk by ID withour ID', []], 
    ['findByIdDisk', 'Find Disk by ID: 1', [1]],
    ['findByIdDisk', 'Find Disk by ID: 2', [2]],
    ['findByIdDisk', 'Find Disk by ID: 3', [3]],
    ['toIncreaseVolume', 'Increase machines number 1 space by disk 1', [{'id' : 1} ,{'id' : 1}]],
    ['toIncreaseVolume', 'Increase machines number 1 space by disk 2', [{'id' : 1} ,{'id' : 2}]],
    ['toIncreaseVolume', 'Increase machines number 1 space by disk 3', [{'id' : 1} ,{'id' : 3}]],
    ['toIncreaseVolume', 'Increase machines number 2 space by disk 1', [{'id' : 2} ,{'id' : 1}]],
    ['toIncreaseVolume', 'Increase machines number 2 space by disk 2', [{'id' : 2} ,{'id' : 2}]],
    ['toIncreaseVolume', 'Increase machines number 2 space by disk 3', [{'id' : 2} ,{'id' : 3}]],
    ['toIncreaseVolume', 'Increase machines number 3 space by disk 1', [{'id' : 3} ,{'id' : 1}]],
    ['toIncreaseVolume', 'Increase machines number 3 space by disk 2', [{'id' : 3} ,{'id' : 2}]],
    ['toIncreaseVolume', 'Increase machines number 3 space by disk 3', [{'id' : 3} ,{'id' : 3}]]
]



const sendTestResponses = async () => {
  const separator = '\n=========================================================\n';
  for (const [ method, comment, args ] of testData)
    try {
      console.log(comment, '\n');
      const responseFn = machineClient[method];
      const res = await responseFn(...args);
      console.log('Result:')
      console.dir(res, { depth: null })
      console.log(separator)
    } catch (e) {
      console.error('Error:', e, '\n', separator);
    }
};

sendTestResponses();

