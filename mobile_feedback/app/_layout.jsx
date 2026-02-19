import {Slot, Stack, useRouter} from "expo-router"
import { Colors } from "../constants/Colors"
import { useRoute } from '@react-navigation/native';
import * as Linking from "expo-linking";
import {useEffect, useState} from "react";
import { UserProvider } from "../contexts/UserContext"


export default function RootLayout() {

  const theme =  Colors.light
  const url = Linking.useURL();
  const [userId,setUserId] = useState("")
    const router = useRouter()

  useEffect(() => {
    console.log('useEffect root:', url);
    if (url) {
      const valueFromUrl = url.split('id=');
      const [urlLink, userId] = valueFromUrl;
      if (typeof userId === 'undefined') {
        handleUrlLink(userId);
      }
        setUserId(userId);
    }
  }, [url]);

    const handleUrlLink = ({ userId }) => {
        if (!userId) {
           router.replace("/home")
        }
    };

  return (
    <UserProvider value={{userId}}>
      <Stack screenOptions={{
        headerStyle: {backgroundColor: theme.navBackground},
        headerTintColor: theme.title,
        headerShown: false,
      }}>
        <Stack.Screen name="(auth)"/>
        <Stack.Screen name="(home)"/>
      </Stack>
    </UserProvider>

  )
}
