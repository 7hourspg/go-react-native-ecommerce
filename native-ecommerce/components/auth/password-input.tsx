import { Button } from '@/components/ui/button';
import { Icon } from '@/components/ui/icon';
import { EyeIcon, EyeOffIcon, LockIcon } from 'lucide-react-native';
import { useState } from 'react';
import { IconInput } from './icon-input';

interface PasswordInputProps {
  label: string;
  placeholder: string;
  value: string;
  onChangeText: (text: string) => void;
  onBlur?: () => void;
  showPassword?: boolean;
  onToggleShowPassword?: () => void;
}

export function PasswordInput({
  label,
  placeholder,
  value,
  onChangeText,
  onBlur,
  showPassword: controlledShowPassword,
  onToggleShowPassword,
}: PasswordInputProps) {
  const [internalShowPassword, setInternalShowPassword] = useState(false);
  const showPassword = controlledShowPassword !== undefined ? controlledShowPassword : internalShowPassword;
  const toggleShowPassword = onToggleShowPassword || (() => setInternalShowPassword(!internalShowPassword));

  return (
    <IconInput
      label={label}
      icon={LockIcon}
      placeholder={placeholder}
      value={value}
      onChangeText={onChangeText}
      onBlur={onBlur}
      secureTextEntry={!showPassword}
      autoCapitalize="none"
      rightElement={
        <Button
          variant="ghost"
          size="icon"
          className="h-8 w-8"
          onPress={toggleShowPassword}>
          <Icon
            as={showPassword ? EyeOffIcon : EyeIcon}
            size={18}
            className="text-muted-foreground"
          />
        </Button>
      }
    />
  );
}

