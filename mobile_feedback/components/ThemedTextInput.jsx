import { TextInput, useColorScheme } from 'react-native'
import { Colors } from '../constants/Colors'

export default function ThemedTextInput({ style, ...props }) {
    const colorScheme = useColorScheme()
    const theme = Colors[colorScheme] ?? Colors.light

    return (
        <TextInput
            style={[
                {
                    backgroundColor: theme.background,
                    borderWidth: 1,           // Required for border to show
                    borderColor: '#d6d5e1',
                    color: theme.text,
                    padding: 18,
                    borderRadius: 15,
                },
                style
            ]}
            {...props}
        />
    )
}
