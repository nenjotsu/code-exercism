namespace targets {
// TODO: Insert the code for the alien class here
  class Alien {
    public:
      int x_coordinate{0};
      int y_coordinate{0};
      Alien(int x, int y){
        x_coordinate = x;
        y_coordinate = y;
      }
      int get_health() {
        return health;
      }
      bool is_alive() {
        return health > 0;
      }
      void hit() {
        health--;
      }
      void teleport(int x, int y) {
        x_coordinate = x;
        y_coordinate = y;
      }
      bool collision_detection(Alien a) {
        return a.x_coordinate == x_coordinate && a.y_coordinate == y_coordinate;
      }


    private:
      int health{3};
  };

}  // namespace targets

int main() {
  return 0;
}