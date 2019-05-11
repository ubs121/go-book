#import('dart:html');

class FallingBall {
  Document doc;
  CanvasElement canvas;
  CanvasRenderingContext2D ctx;
  bool falling;
  int x, y;
  final int width = 320;
  final int height = 150;
  final String ballFile = "ball.png";
  final String readyMsg = "Dart ready.";
  final String fallingMsg = "Falling ...";
  ImageElement img;
  final double oAccel = .45;
  final double oXSpeed = 1.3;
  double accel;
  double xSpeed;
  double t = 0.0;
  num lastRendered = null;
  bool stopped;
  bool xDir;
  
  FallingBall() {}
  
  void go() {
    setupCanvas();
    
    document.query("#status").innerHTML = readyMsg;
    document.query("#fall").on.click.add((event) { xDir = true; x=0; y=0; init(); window.requestAnimationFrame(fall); });
    document.query("#canvas").on.click.add((event) => kickStart());
  }
  
  void updateStatus() {
    if (!stopped)
      document.query("#status").innerHTML = fallingMsg;
    else
      document.query("#status").innerHTML = readyMsg;
  }
  
  void setupCanvas() {
    doc = window.document;
    canvas = doc.query("#canvas");
    canvas.width = width;
    canvas.height = height;
    ctx = canvas.getContext("2d");
    loadImage(ballFile);
  }
  
  void loadImage(String src) {
    img = new Element.tag("img");
    img.src = src;
  }
  
  void displayBall(ImageElement img, int x, int y) {
    ctx.drawImage(img, x, y);
  }
  
  void init() {
    falling = true;
    stopped = false;
    lastRendered = (new Date.now()).milliseconds;
    t = 0.0;
    accel = oAccel;
    xSpeed = oXSpeed;
  }
  
  void fall(int time) {
    var delta = (new Date.now()).milliseconds - lastRendered;

    if (falling)
      t += accel;
    else
      t -= accel;
    ctx.clearRect(0, 0, width, height);
    ctx.setFillColor("#dddddd");
    ctx.fillRect(0, height-1, width, height);
        
    if (y + img.height < height - 1 && falling) {
        y += delta / 1000 + t;
        if (y + img.height >= height - 1) {
          falling = false;
          accel += .25 * y / height;
        }
    }
    else if (y > 0 && !falling) {
      y -= delta / 1000 + t;
      if (y <= 0 || t <= 0) {
        falling = true;
        t = 0.0;
      }
    }
    else if (y < 0 && !falling) {
      falling = true;
      accel += .25 * y / height;
    }
    else {
      xSpeed += .018 * (xDir ? -1 : 1);
      if ((xSpeed < 0 && xDir) || (xSpeed > 0 && !xDir))
        stopped = true;
    }
    
    if (y + img.height >= height - 1)
      y = height - img.height - 1;
    
    if (x < 0 || x + img.width > width) {
      xDir = !xDir;
      xSpeed = -xSpeed;
    }

    displayBall(img, x += xSpeed, y);
    updateStatus();
    document.query("#status").innerHTML;
    if (!stopped) {
      document.query("#fall").attributes.putIfAbsent("disabled", () => "disabled");
      window.requestAnimationFrame(fall);
    }
    else {
      document.query("#fall").attributes.remove("disabled");
      init();
      stopped = true;
      updateStatus();
    }
  }  
  
  void kickStart() {
    if (stopped) {
      bool xd = xDir;
      init();
      xDir = xd;
      xSpeed = xSpeed * (xDir ? 1 : -1);
      falling = false;
      y--;
      t = Math.random() * 11 + 1;
      window.requestAnimationFrame(fall);
    }
  }
  
}

void main() {
  FallingBall fb = new FallingBall();
  fb.go();
}