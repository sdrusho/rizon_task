import { Image, useColorScheme } from 'react-native'

// images
import DarkLogo from '../assets/images/rizon_logo.jpeg'
import LightLogo from '../assets/images/rizon_logo.jpeg'

const ThemedLogo = () => {
    const colorScheme = useColorScheme()

    const logo = colorScheme === 'dark' ? DarkLogo : LightLogo

    return (
        <Image style={{borderRadius: 20,height: 120, width: 120}} source={logo} />
    )
}

export default ThemedLogo
