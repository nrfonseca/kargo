import {
  faCircle,
  faCircleNotch,
  faHeart,
  faHeartBroken,
  faQuestionCircle,
  IconDefinition
} from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { Tooltip } from 'antd';
import { CSSProperties } from 'react';

import { Health } from '@ui/gen/v1alpha1/types_pb';

import { HealthStatus, healthStatusToEnum } from './utils';

export const HealthStatusIcon = (props: {
  health?: Health;
  style?: CSSProperties;
  hideColor?: boolean;
}) => {
  const { health, hideColor } = props;
  const reason = health?.issues?.join('; ') ?? '';

  return (
    <Tooltip title={health?.status ?? '' + (reason !== '' ? `: ${reason}` : '')}>
      <FontAwesomeIcon
        icon={iconForHealthStatus(health)}
        spin={healthStatusToEnum(health?.status) === HealthStatus.PROGRESSING}
        style={{
          color: !hideColor ? colorForHealthStatus(health) : undefined,
          fontSize: '18px',
          ...props.style
        }}
      />
    </Tooltip>
  );
};

const iconForHealthStatus = (health?: Health): IconDefinition => {
  switch (healthStatusToEnum(health?.status)) {
    case HealthStatus.HEALTHY:
      return faHeart;
    case HealthStatus.UNHEALTHY:
      return faHeartBroken;
    case HealthStatus.PROGRESSING:
      return faCircleNotch;
    case HealthStatus.UNKNOWN:
      return faQuestionCircle;
    default:
      return faCircle;
  }
};

const colorForHealthStatus = (health?: Health): string => {
  switch (healthStatusToEnum(health?.status)) {
    case HealthStatus.HEALTHY:
      return '#52c41a';
    case HealthStatus.UNHEALTHY:
      return '#f5222d';
    case HealthStatus.PROGRESSING:
      return '#0dabea';
    case HealthStatus.UNKNOWN:
      return '#faad14';
    default:
      return '#ccc';
  }
};
