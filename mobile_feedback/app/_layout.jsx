import {Slot, Stack, useRouter} from "expo-router"
import { Colors } from "../constants/Colors"
import { useRoute } from '@react-navigation/native';
import * as Linking from "expo-linking";
import {useEffect, useState} from "react";
import { UserProvider } from "../contexts/UserContext"
import {useUser} from "../hooks/useUser";


export default function RootLayout() {

  const theme =  Colors.light
  const url = Linking.useURL();
  const [userId,setUserId] = useState("")
  useEffect(() => {
    console.log('useEffect root:', url);
    if (url) {
      const valueFromUrl = url.split('id=');
      const [urlLink, userId]= valueFromUrl;
      console.log('App opened with URL from root:', userId);
      setUserId(userId)

    }
  }, [url]);

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
