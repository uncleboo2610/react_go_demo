import './App.css';
import { Routes, Route } from 'react-router-dom';
import Posts from './components/posts/Posts';
import PostForm from './components/posts/PostForm';

function App() {
  return (
    <Routes>
      <Route path='/posts' element={<Posts />} />
      <Route path='/postForm' element={<PostForm />} />
    </Routes>
  );
}

export default App;
