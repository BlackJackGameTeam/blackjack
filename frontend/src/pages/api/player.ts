import { CsrfToken } from "@/types";
import axios from "axios";

export const getCsrfToken = async () => {
    const { data } = await axios.get<CsrfToken>(
        `${process.env.NEXT_PUBLIC_API_URL}/csrf`
    )
    axios.defaults.headers.common['X-CSRF-Token'] = data.csrf_token
}