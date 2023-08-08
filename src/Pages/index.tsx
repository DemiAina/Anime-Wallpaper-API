import { Link } from "react-router-dom"
const Index = () =>{
    return (
        <div>
            <div>
                <img id = "hero_image" src="/71XtBFClyOL.jpg"/>
            </div>
            <ul id = "upload">
            <li>
                Upload
            </li>
            <li>
                <Link to= "#">View images</Link>
            </li>
            </ul>
        </div>
    )
}

export default Index
