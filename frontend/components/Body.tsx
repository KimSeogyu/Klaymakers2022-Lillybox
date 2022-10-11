import Navbar from "./Navbar";
import VideoPlayer from "./VideoPlayer";

function Body() {
  return (
    <div className="col-span-6 relative min-h-screen z-0">
      <div className="ml-4 lg:ml-17">
        <Navbar />
        <VideoPlayer />
      </div>
    </div>
  );
}

export default Body;
