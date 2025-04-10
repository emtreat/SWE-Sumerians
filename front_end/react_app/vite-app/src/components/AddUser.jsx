import '../App.css'
import { PostUser } from './PostUser'

export function AddUser() {
    return (
        <button className="floating-button">
        <i className="material-icons">add</i> 
        <Link to="/Home"> {<PostUser />}</Link>
      </button>
    )
}