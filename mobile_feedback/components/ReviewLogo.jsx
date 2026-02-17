import { Image, useColorScheme } from 'react-native'


import reviewlogo from '../assets/images/review_logo.png'


const ReviewLogo = () => {
    const colorScheme = useColorScheme()

    const logo = reviewlogo

    return (
        <Image style={{borderRadius: 20,height: 150, width: 150}} source={logo} />
    )
}

export default ReviewLogo
