import { createContext, useState } from "react"
import { useRouter } from 'expo-router'



export const UserContext = createContext()
export const API_URL = "http://192.168.110.11:8001/ms-feedback";


export function UserProvider({ children }) {
    const router = useRouter()
    const [userId, setUserId] = useState(null)
    const [isEnjoying, setIsEnjoying] = useState(false)
    const [comments, setComments] = useState("")
    const [review, setReview] = useState(false)
    const [token, setTokenId] = useState("")
    const [ user, setUser] = useState(null)


    async function login(userId) {
        try {
            const getUrl = `${API_URL}/users/${userId}`;
            const response = await fetch(getUrl)
            const userData = await response.json()
            setUser(userData)
            setUserId(userId)
            if (!response.ok) {
                const errorData = await response.json();
                console.log(errorData);
                router.replace("/fail")
            }
        } catch (error) {
            console.log("error from login")
            router.replace("/index")
        }
    }

    async function sendIsEnjoying(userId, enjoying) {
        try {
            setIsEnjoying(enjoying)
            setUserId(userId)
        } catch (error) {
            throw Error(error.message)
        }
    }

    async function sendFeedBackComments(userId, comments) {
        try {
           setIsEnjoying(isEnjoying)
            setComments(comments)
        } catch (error) {
            throw Error(error.message)
        }
    }

    async function sendReview(userId, review) {
        setReview(review)
        try {
            const token = user.accessToken
            const id = userId ?? user.userId
            const response = await fetch(`${API_URL}/feedbacks`, {
                method: "POST",
                headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`
                },
                body: JSON.stringify({
                    "isEnjoying": isEnjoying,
                    "isLeaveReview": review,
                    "comments": comments,
                    "userId": id,
                }),
            });
            //openStore();
           response.ok ? router.replace("/successSignup") : router.replace("/userNotFound")
        } catch (error) {
            console.log(error.message);
            router.replace("/userNotFound")
        }

    }



    return (
        <UserContext.Provider value={{
            login,sendIsEnjoying,sendFeedBackComments,sendReview,userId,
        }}>
            {children}
        </UserContext.Provider>
    );
}

const GOOGLE_PACKAGE_NAME = 'com.example.myapp'; // Your Android Package Name
const APPLE_APP_ID = '123456789'; // Your iOS App ID (found in App Store Connect)

const openStore = () => {
    const url = Platform.select({
        ios: `itms-apps://itunes.apple.com/app/id${APPLE_APP_ID}`,
        android: `market://details?id=${GOOGLE_PACKAGE_NAME}`,
    });

    Linking.canOpenURL(url)
        .then((supported) => {
            if (supported) {
                Linking.openURL(url);
            } else {
                // Fallback to browser links if store apps aren't available (e.g., on emulators)
                const browserUrl = Platform.select({
                    ios: `https://apps.apple.com/app/id${APPLE_APP_ID}`,
                    android: `https://play.google.com/store/apps/details?id=${GOOGLE_PACKAGE_NAME}`,
                });
                Linking.openURL(browserUrl);
            }
        })
        .catch((err) => console.error('An error occurred', err));
};

