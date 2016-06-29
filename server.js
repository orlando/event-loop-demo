var messages = [
  { color: '#00d2ff', index: 1 },
  { color: '#3a7bd5', index: 2 },
  { color: '#D3959B', index: 3 },
  { color: '#BFE6BA', index: 4 },
  { color: '#DAD299', index: 5 },
  { color: '#B0DAB9', index: 6 },
  { color: '#E6DADA', index: 7 },
  { color: '#274046', index: 8 },
  { color: '#5D4157', index: 9 },
  { color: '#A8CABA', index: 10 },
  { color: '#ddd6f3', index: 11 },
  { color: '#faaca8', index: 12 },
  { color: '#616161', index: 13 },
  { color: '#9bc5c3', index: 14 },
  { color: '#50C9C3', index: 15 },
  { color: '#96DEDA', index: 16 },

  { color: '#00d2ff', index: 17 },
  { color: '#3a7bd5', index: 18 },
  { color: '#D3959B', index: 19 },
  { color: '#BFE6BA', index: 20 },
  { color: '#DAD299', index: 21 },
  { color: '#B0DAB9', index: 22 },
  { color: '#E6DADA', index: 23 },
  { color: '#274046', index: 24 },
  { color: '#5D4157', index: 25 },
  { color: '#A8CABA', index: 26 },
  { color: '#ddd6f3', index: 27 },
  { color: '#faaca8', index: 28 },
  { color: '#616161', index: 29 },
  { color: '#9bc5c3', index: 30 },
  { color: '#50C9C3', index: 31 },
  { color: '#96DEDA', index: 32 },

  { color: '#00d2ff', index: 33 },
  { color: '#3a7bd5', index: 34 },
  { color: '#D3959B', index: 35 },
  { color: '#BFE6BA', index: 36 },
  { color: '#DAD299', index: 37 },
  { color: '#B0DAB9', index: 38 },
  { color: '#E6DADA', index: 39 },
  { color: '#274046', index: 40 },
  { color: '#5D4157', index: 41 },
  { color: '#A8CABA', index: 42 },
  { color: '#ddd6f3', index: 43 },
  { color: '#faaca8', index: 44 },
  { color: '#616161', index: 45 },
  { color: '#9bc5c3', index: 46 },
  { color: '#50C9C3', index: 47 },
  { color: '#96DEDA', index: 48 },

  { color: '#00d2ff', index: 49 },
  { color: '#3a7bd5', index: 50 },
  { color: '#D3959B', index: 51 },
  { color: '#BFE6BA', index: 52 },
  { color: '#DAD299', index: 53 },
  { color: '#B0DAB9', index: 54 },
  { color: '#E6DADA', index: 55 },
  { color: '#274046', index: 56 },
  { color: '#5D4157', index: 57 },
  { color: '#A8CABA', index: 58 },
  { color: '#ddd6f3', index: 59 },
  { color: '#faaca8', index: 60 },
  { color: '#616161', index: 61 },
  { color: '#9bc5c3', index: 62 },
  { color: '#50C9C3', index: 63 },
  { color: '#96DEDA', index: 64 },
];

var messagesShort = messages.slice(0,6)

var io = require('socket.io')(8080);

io.on('connection', function (socket) {
  console.log('new connection!');
  var i = 0;

  socket.on('ready', function () {
    messagesShort.forEach(function (message) {
      console.log('sending message', message);

      socket.emit('paint', message);

      while (i < 1000000000) {
        i++
      }

      i = 0;
    });
  });
});

console.log('running!');
