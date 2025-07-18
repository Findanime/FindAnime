
import SearchPane from "../../components/home/SearchPane"   
import About from "../../components/home/About"
export function Home(){
    return (
        <>
        <SearchPane/>
        <About positionStyle={{ top: '800px', left: '0px', right: '50px' , maxWidth : '1000px'} }/>
        </>
    )
}