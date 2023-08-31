import React from "react";
import ReactDOM from "react-dom";

import api from "./api"

function App() {
  const [likesCounter, setLikesCounter] = React.useState(0)

  React.useEffect(() => {
    const fetchLikes = async () => {
      setLikesCounter(await api.LikesCount())
    }
    fetchLikes()
      .catch(console.error)
  }, [likesCounter]);

  const onLikeClick = async () => {
    await api.LikesIncrement()
    setLikesCounter(prevCount => prevCount + 1)
  }

  return (
    <div className="mx-auto flex flex-1 flex-col items-center justify-center">
      <button
        className="relative inline-block text-lg group"
        // onClick={() => setLikesCounter(prevCount => prevCount + 1)}
        onClick={onLikeClick}
      >
        <span className="relative z-10 block px-5 py-3 overflow-hidden font-medium leading-tight text-gray-800 transition-colors duration-300 ease-out border-2 border-gray-900 rounded-lg group-hover:text-white">
          <span className="absolute inset-0 w-full h-full px-5 py-3 rounded-lg bg-gray-50"></span>
          <span className="absolute left-0 w-48 h-48 -ml-2 transition-all duration-300 origin-top-right -rotate-90 -translate-x-full translate-y-12 bg-gray-900 group-hover:-rotate-180 ease"></span>
          <span className="relative">Like It!</span>
        </span>
        <span
          className="absolute bottom-0 right-0 w-full h-12 -mb-1 -mr-1 transition-all duration-200 ease-linear bg-gray-900 rounded-lg group-hover:mb-0 group-hover:mr-0"
          data-rounded="rounded-lg"
        ></span>
      </button>
      <h1 className="mt-2">{likesCounter} people liked this page</h1>
    </div>
  );
}

const root = ReactDOM.createRoot(document.getElementById("app"));
root.render(<App />);
