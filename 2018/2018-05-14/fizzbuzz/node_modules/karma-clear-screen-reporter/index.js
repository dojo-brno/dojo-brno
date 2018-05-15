function ClearScreenReporter() {
  this.onRunStart = function() {
    console.log("\u001b[2J\u001b[0;0H");
  };
}

module.exports = {
  'reporter:clear-screen': ['type', ClearScreenReporter]
};
