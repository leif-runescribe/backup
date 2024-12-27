import axios from "axios";
import { useEffect, useState } from "react";

const base_url = "https://api.diadata.org/v1/quotation/"

const useDia = ()=>{
    const [price, setprice] = useState(null)
    const [loading, setloading] = useState(true)
    const [error, seterror] = useState(null)

    useEffect(()=>{
        const fetch = async()=>{
            try {
                const respe = await axios.get(`${base_url}/usd`);
                setprice({respe.data})
                console.log(respe)
            } catch (error) {
                console.log(error)
            }
            finally{
                setloading(false)
            }
        }
        fetch
    })
}

export default useDia