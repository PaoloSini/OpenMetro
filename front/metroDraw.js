

// first we need to create a stage
var stage = new Konva.Stage({
  container: 'canvas',   // id of container <div>
  width: 1920,
  height: 1080
});

// then create layer
var layer = new Konva.Layer();

function drawCircle(posX, posY, color="red") {
  // create our shape
  return new Konva.Circle({
    x: posX,
    y: posY,
    radius: 10,
    fill: color,
    stroke: 'black',
    strokeWidth: 4
  });
}

function drawSquare(posX, posY, color='blue') {
  return new Konva.Rect({
    x: posX,
    y: posY,
    height: 15,
    width: 15,
    fill: color,
  })
}

function drawLine(points, color) {

  return new Konva.Line({
    points: points,
    stroke: color,
    strokeWidth: 5,
  })
}

// add the layer to the stage
stage.add(layer);

function updateStations(stations) {
  Object.entries(stations).forEach(entry => {
    const [stationName, station] = entry;
    layer.add(drawCircle(station.PosX, station.PosY))
  });
}

function updateLines(lines, stations) {
  Object.entries(lines).forEach(entry => {
    const [lineName, line] = entry;
    let linePoints = []
    line.Stations.forEach( station => {
      linePoints.push(station.PosX, station.PosY)
    })
    layer.add(drawLine(linePoints, line.Color))
  });
}

function updateTrains(trains) {
  trains.forEach(train => {
    layer.add(drawSquare(train.PosX, train.PosY, 'blue'))
  })
}

export function updateMap(metroMap) {
  layer.destroyChildren()
  updateStations(metroMap.Stations)
  updateLines(metroMap.Lines, metroMap.Stations)
  updateTrains(metroMap.Trains)
  layer.draw();
}
