import '../App.css'
import { PostUser } from './PostUser'
import { Post } from '../pages/AddPage'
import { AddFile } from '../pages/UploadFile'

export function AddUser() {
    return (
        <button className="floating-button">
        <i className="material-icons">Add File</i> 
      </button>
    )
}