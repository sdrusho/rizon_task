import { Text } from "react-native";
import { StyleSheet } from 'react-native'
import { useState, useEffect } from 'react'

import ThemedLogo from "../components/ThemedLogo";
import ThemedView from '../components/ThemedView'
import ThemedText from '../components/ThemedText'
import Spacer from '../components/Spacer'
import ThemedButton from '../components/ThemedButton'
import ThemedTextInput from "../components/ThemedTextInput"
import {Link, useRouter} from "expo-router";


import { Colors } from '../constants/Colors'

export const SIGNUP_URL = "http://192.168.110.11:8001/ms-feedback/user-signup";
export default function Index() {
  const [uniqueId, setUniqueId] = useState('');
  const [tokenId, setTokenId] = useState('')

  const [email, setEmail] = useState("")
  const router = useRouter()
  //const { user, login } = useUser()
  const [error, setError] = useState()
  /*useEffect(() => {
      const fetchUniqueId = async () => {
          const id = await DeviceInfo.getUniqueId();
          setUniqueId(id);
      };

      fetchUniqueId();
  }, []);*/

  const handleSignup = async () => {
    setError(null)
    try {
      const response = await fetch(`${SIGNUP_URL}`, {
        method: "POST",
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          "email": email,
          "name": "rusho",
          "deviceId": "device001",
        }),
      });
      if (!response.ok) {
        const errorData = await response.json();
        console.log(errorData);
        throw new Error(errorData.error || "Failed to sing up");
        router.replace("/userNotFound")
      }
      router.replace("/success")
    } catch (error) {
      setError(error.message)
    }
  }

  return (

    <ThemedView style={styles.container}>
      <ThemedLogo />
      <Spacer />

      {/* <TextInput placeholder="Email" /> */}

      <Spacer />
      <ThemedTextInput
        style={{ marginBottom: 20, width: "80%" }}
        placeholder="Email"
        value={email}
        onChangeText={setEmail}
        keyboardType="email-address"
      />

      <ThemedButton style={{marginHorizontal: 3,width: '80%'}} onPress={handleSignup}>
        <Text style={{ color: '#f2f2f2',alignSelf: "center", }}>Signup</Text>
      </ThemedButton>

      <Spacer />
      {error && <Text style={styles.error}>{error}</Text>}
    </ThemedView>


  )
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'center',
  },
  img: {
    marginVertical: 20
  },
  title: {
    fontWeight: 'bold',
    fontSize: 18,
  },
  link: {
    marginVertical: 10,
    borderBottomWidth: 1
  },
  error: {
    color: Colors.warning,
    padding: 10,
    backgroundColor: '#f5c1c8',
    borderColor: Colors.warning,
    borderWidth: 1,
    borderRadius: 6,
    marginHorizontal: 10,
  },
})
