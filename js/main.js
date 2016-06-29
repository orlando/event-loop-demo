var UIMapper = function UIMapper(element) {
};

UIMapper.element = document.getElementById('ui');

UIMapper.paint = function paint(data) {
  var index = data.index - 1;
  var row = Math.floor(index / 8);
  var column = index % 8;
  var color = data.color;

  var rowElement = this.element.children[row];
  var columnElement = rowElement.children[column];

  return columnElement.style.background = color;
};

var init = function init() {
  window.socket = io('http://localhost:8080/');

  window.socket.on('paint', function paint(data) {
    console.log('receiving: ' + data);
    UIMapper.paint(data);
  });

  window.socket.emit('ready');
};

window.addEventListener('load', init);
